package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"cedarhawk/internal/ai"
	"cedarhawk/internal/config"
	"cedarhawk/internal/crawler"
	"cedarhawk/internal/results"
	"cedarhawk/internal/ui"
)

// customUsage defines the help message displayed when --help is used.
func customUsage() {
	usageText := `
CedarHawk - Automated Website Functionality Testing Tool

Usage:
  cedarhawk [options]

Options:
  --config       Path to configuration file (default: configs/config.toml)
  --crawl-only   Run only the crawler module
  --ui-only      Run only the UI tests
  --help         Display this help message

Description:
  CedarHawk tests the functionality of a website by performing crawling, UI tests (layout, elements, and responsiveness),
  and optionally invoking an AI analysis plugin. The tool outputs a consolidated JSON report summarizing all test results.

For detailed documentation, please refer to the /docs folder.
`
	fmt.Fprintf(os.Stderr, "%s\n", usageText)
}

func main() {
	// Override the default usage function.
	flag.Usage = customUsage

	// Define command-line flags.
	configFile := flag.String("config", "configs/config.toml", "Path to configuration file")
	crawlOnly := flag.Bool("crawl-only", false, "Run only the crawler module")
	uiOnly := flag.Bool("ui-only", false, "Run only the UI tests")
	flag.Parse()

	// If help flag is provided, display usage and exit.
	if flag.NArg() == 0 && (len(os.Args) > 1 && os.Args[1] == "--help") {
		flag.Usage()
		os.Exit(0)
	}

	// Load configuration.
	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize final results structure.
	finalResults := results.TestResults{
		Page:                  cfg.StartURL,
		ErrorLogs:             []results.ErrorLog{},
		LayoutResults:         []results.LayoutResult{},
		UIElementsResults:     []results.UIElementResult{},
		ResponsivenessResults: []results.ResponsivenessResult{},
	}

	// Helper function to record errors.
	recordError := func(module, message string) {
		finalResults.ErrorLogs = append(finalResults.ErrorLogs, results.ErrorLog{
			Module:    module,
			Message:   message,
			Timestamp: time.Now().Format(time.RFC3339),
		})
	}

	// Variable to hold crawler output.
	var pages []crawler.PageData

	// Run Crawler module if --crawl-only is set or if running full suite.
	if *crawlOnly || (!*crawlOnly && !*uiOnly) {
		log.Println("Running crawler module...")
		var err error
		pages, err = crawler.Crawl(cfg.StartURL, cfg.MaxDepth)
		if err != nil {
			msg := "Crawler error: " + err.Error()
			log.Println(msg)
			recordError("Crawler", msg)
		}
		if len(pages) == 0 {
			msg := "Crawler returned no pages."
			log.Println(msg)
			recordError("Crawler", msg)
		}
	}

	// Run UI tests if --ui-only is set or if running full suite.
	if *uiOnly || (!*crawlOnly && !*uiOnly) {
		log.Println("Running UI tests...")

		// Use the HTML from the first crawled page, if available.
		sampleHTML := ""
		if len(pages) > 0 {
			sampleHTML = pages[0].HTML
		} else {
			log.Println("No crawled HTML available; UI tests will run on empty input.")
		}

		// Layout Verification.
		layoutResult := ui.RunLayoutVerification(sampleHTML, cfg.UI.ExpectedLayout)
		finalResults.LayoutResults = append(finalResults.LayoutResults, layoutResult)

		// UI Elements Validation.
		elementResults := ui.RunElementsValidation(sampleHTML)
		finalResults.UIElementsResults = elementResults

		// Responsiveness Testing.
		responsiveResults := ui.RunResponsivenessTests(sampleHTML, cfg.UI.Viewports)
		finalResults.ResponsivenessResults = responsiveResults
	}

	// Invoke AI plugin.
	aiPlugin := &ai.DefaultAIPluginStub{}
	aiReport, err := aiPlugin.AnalyzeResults(&finalResults)
	if err != nil {
		msg := "AI Plugin error: " + err.Error()
		log.Println(msg)
		recordError("AI Plugin", msg)
	} else {
		finalResults.AIAnalysis = aiReport
	}

	// Write consolidated JSON output.
	if err := results.WriteJSON(finalResults, cfg.ResultsFile); err != nil {
		log.Fatalf("Failed to write JSON results: %v", err)
	}
	log.Printf("CedarHawk completed. Results written to %s\n", cfg.ResultsFile)
	os.Exit(0)
}

