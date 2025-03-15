package integration

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"cedarhawk/internal/ai"
	"cedarhawk/internal/config"
	"cedarhawk/internal/crawler"
	"cedarhawk/internal/logging"
	"cedarhawk/internal/results"
	"cedarhawk/internal/ui"
)

// replace is a simple helper function to replace substrings.
func replace(input, old, new string) string {
	return strings.ReplaceAll(input, old, new)
}

func TestEndToEndWorkflow(t *testing.T) {
	// --- Step 1: Create a temporary configuration file.
	configContent := `
startURL = "http://127.0.0.1:0/testpage"
maxDepth = 1
resultsFile = "test_results.json"

[ui]
expectedLayout = "default"
[[ui.viewports]]
name = "desktop"
width = 1920
height = 1080
[[ui.viewports]]
name = "mobile"
width = 375
height = 667

[logging]
level = "info"
`
	tmpConfigFile, err := ioutil.TempFile("", "config-*.toml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpConfigFile.Name())
	if _, err := tmpConfigFile.Write([]byte(configContent)); err != nil {
		t.Fatal(err)
	}
	tmpConfigFile.Close()

	// --- Step 2: Set up a test HTTP server to simulate a website.
	mux := http.NewServeMux()
	mux.HandleFunc("/testpage", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
				<header>Header</header>
				<nav>Navigation</nav>
				<main>Main Content</main>
				<footer>Footer</footer>
				<button onclick="alert('clicked')">Click me</button>
				<a href="http://example.com">Example Link</a>
				<img src="` + r.Host + `/image.png" alt="Test Image">
				<input type="text" name="username">
				<textarea name="comments"></textarea>
				<div class="desktop-responsive">Desktop Layout</div>
				<div class="mobile-responsive">Mobile Layout</div>
			</html>
		`))
	})
	mux.HandleFunc("/image.png", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("PNG DATA"))
	})
	testServer := httptest.NewServer(mux)
	defer testServer.Close()

	// Update the configuration file's startURL with the actual test server URL.
	configData, err := ioutil.ReadFile(tmpConfigFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	updatedConfig := replace(string(configData), "http://127.0.0.1:0/testpage", testServer.URL+"/testpage")
	if err := ioutil.WriteFile(tmpConfigFile.Name(), []byte(updatedConfig), 0644); err != nil {
		t.Fatal(err)
	}

	// --- Step 3: Load configuration.
	cfg, err := config.LoadConfig(tmpConfigFile.Name())
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// --- Step 4: Initialize logging.
	logging.InitializeLogger()

	// --- Step 5: Run the crawler.
	pages, err := crawler.Crawl(cfg.StartURL, cfg.MaxDepth)
	if err != nil {
		t.Fatalf("Crawler error: %v", err)
	}
	if len(pages) == 0 {
		t.Fatal("Crawler returned no pages")
	}

	// --- Step 6: Run UI tests using the HTML from the first crawled page.
	sampleHTML := pages[0].HTML
	layoutResult := ui.RunLayoutVerification(sampleHTML, cfg.UI.ExpectedLayout)
	elementResults := ui.RunElementsValidation(sampleHTML)
	responsiveResults := ui.RunResponsivenessTests(sampleHTML, cfg.UI.Viewports)

	// --- Step 7: Build final test results structure.
	finalResults := results.TestResults{
		Page:                  cfg.StartURL,
		LayoutResults:         []results.LayoutResult{layoutResult},
		UIElementsResults:     elementResults,
		ResponsivenessResults: responsiveResults,
		ErrorLogs:             logging.GetCapturedErrors(),
	}

	// --- Step 8: Invoke AI plugin.
	aiPlugin := &ai.DefaultAIPluginStub{}
	aiReport, err := aiPlugin.AnalyzeResults(&finalResults)
	if err != nil {
		t.Fatalf("AI Plugin error: %v", err)
	}
	finalResults.AIAnalysis = aiReport

	// --- Step 9: Write the consolidated JSON output.
	resultsFile := cfg.ResultsFile
	err = results.WriteJSON(finalResults, resultsFile)
	if err != nil {
		t.Fatalf("Failed to write JSON results: %v", err)
	}

	// --- Step 10: Read and verify the JSON results.
	data, err := ioutil.ReadFile(resultsFile)
	if err != nil {
		t.Fatalf("Failed to read JSON results file: %v", err)
	}
	var parsed results.TestResults
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("Failed to unmarshal JSON results: %v", err)
	}
	// Check that layout results and UI element results exist.
	if len(parsed.LayoutResults) == 0 {
		t.Error("Expected at least one layout result, got none")
	}
	if len(parsed.UIElementsResults) == 0 {
		t.Error("Expected UI elements results, got none")
	}
	// Clean up the JSON results file.
	os.Remove(resultsFile)
}

