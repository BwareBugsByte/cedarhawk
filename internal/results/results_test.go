package results

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	// Create a sample TestResults structure.
	sampleResults := TestResults{
		Page: "https://example.com",
		LayoutResults: []LayoutResult{
			{Page: "https://example.com", Status: "pass", Message: "Layout is consistent."},
		},
		UIElementsResults: []UIElementResult{
			{ElementType: "button", Status: "pass", Message: "Button is present.", ElementIdentifier: "btn-submit"},
			{ElementType: "image", Status: "fail", Message: "Image source returned HTTP 404.", ElementIdentifier: "img-logo"},
		},
		ResponsivenessResults: []ResponsivenessResult{
			{Viewport: "desktop", Status: "pass"},
			{Viewport: "mobile", Status: "fail", Issues: []string{"Element clipped", "Navigation menu overlaps"}},
		},
		ErrorLogs: []ErrorLog{
			{Module: "Crawler", Message: "Failed to fetch URL https://example.com/page", Timestamp: "2025-02-18T10:23:45Z"},
		},
		AIAnalysis: &AIAnalysisReport{
			Summary:         "AI analysis not available.",
			Recommendations: []string{}, // explicitly empty slice
		},
	}

	// Create a temporary file for testing.
	tmpFile, err := ioutil.TempFile("", "results-*.json")
	if err != nil {
		t.Fatal(err)
	}
	tmpFileName := tmpFile.Name()
	tmpFile.Close()
	defer os.Remove(tmpFileName)

	// Write JSON using the WriteJSON function.
	if err := WriteJSON(sampleResults, tmpFileName); err != nil {
		t.Fatalf("WriteJSON returned an error: %v", err)
	}

	// Read the file content.
	data, err := ioutil.ReadFile(tmpFileName)
	if err != nil {
		t.Fatalf("Failed to read JSON file: %v", err)
	}

	// Unmarshal the JSON back into a TestResults struct.
	var parsedResults TestResults
	if err := json.Unmarshal(data, &parsedResults); err != nil {
		t.Fatalf("Failed to unmarshal JSON content: %v", err)
	}

	// Normalize Recommendations slices in AIAnalysis (treat nil as empty slice).
	if sampleResults.AIAnalysis != nil {
		if sampleResults.AIAnalysis.Recommendations == nil {
			sampleResults.AIAnalysis.Recommendations = []string{}
		}
	}
	if parsedResults.AIAnalysis != nil {
		if parsedResults.AIAnalysis.Recommendations == nil {
			parsedResults.AIAnalysis.Recommendations = []string{}
		}
	}

	// Compare AIAnalysis values by dereferencing pointers.
	if sampleResults.AIAnalysis != nil && parsedResults.AIAnalysis != nil {
		if !reflect.DeepEqual(*sampleResults.AIAnalysis, *parsedResults.AIAnalysis) {
			t.Errorf("AIAnalysis content mismatch. Expected: %+v, Got: %+v", *sampleResults.AIAnalysis, *parsedResults.AIAnalysis)
		}
		// Set the pointer in sampleResults to match parsedResults for full comparison.
		sampleResults.AIAnalysis = parsedResults.AIAnalysis
	}

	// Verify that the original and parsed results are equal.
	if !reflect.DeepEqual(sampleResults, parsedResults) {
		t.Errorf("Expected:\n%+v\nGot:\n%+v", sampleResults, parsedResults)
	}
}
