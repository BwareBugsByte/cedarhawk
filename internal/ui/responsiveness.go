package ui

import (
	"log"
	"strings"

	"cedarhawk/internal/config"
	"cedarhawk/internal/results"
)

// RunResponsivenessTests simulates responsiveness testing by checking for
// viewport-specific markers in the provided HTML data. For example, for a viewport
// named "mobile", it expects the HTML to contain the marker "mobile-responsive".
// It returns a slice of ResponsivenessResult detailing the status for each viewport.
func RunResponsivenessTests(htmlData string, viewports []config.Viewport) []results.ResponsivenessResult {
	var resultsSlice []results.ResponsivenessResult

	// Convert HTML to lower-case for case-insensitive matching.
	lowerHTML := strings.ToLower(htmlData)

	// Process each viewport configuration.
	for _, vp := range viewports {
		result := results.ResponsivenessResult{
			Viewport: vp.Name,
			Status:   "pass",
			Issues:   []string{},
		}

		// Expected marker format: "<viewport>-responsive" (e.g., "mobile-responsive")
		expectedMarker := strings.ToLower(vp.Name) + "-responsive"
		if !strings.Contains(lowerHTML, expectedMarker) {
			result.Status = "fail"
			issue := "Expected responsive marker '" + expectedMarker + "' not found."
			result.Issues = append(result.Issues, issue)
			log.Printf("Responsiveness Test: For viewport %s, marker '%s' not found.", vp.Name, expectedMarker)
		}

		resultsSlice = append(resultsSlice, result)
	}

	return resultsSlice
}
