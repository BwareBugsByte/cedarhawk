package ui

import (
	"log"
	"strings"

	"cedarhawk/internal/results"
)

// RunLayoutVerification extracts the DOM layout from the provided HTML and compares
// it against the expected layout pattern defined in expectedLayout.
// For the "default" layout, we expect to see a <header>, <nav>, <main>, and <footer> element.
// It returns a LayoutResult that indicates pass/fail status and a descriptive message.
func RunLayoutVerification(htmlData, expectedLayout string) results.LayoutResult {
	// Define expected elements based on the expected layout.
	var expectedElements []string
	if expectedLayout == "default" {
		expectedElements = []string{"<header", "<nav", "<main", "<footer"}
	} else {
		// For custom layouts, you may define other expected elements.
		expectedElements = []string{"<div", "<p"}
	}

	var missingElements []string
	for _, tag := range expectedElements {
		// Using case-insensitive check by converting htmlData to lower case.
		if !strings.Contains(strings.ToLower(htmlData), tag) {
			missingElements = append(missingElements, tag)
			log.Printf("Layout Verification: Missing expected element: %s", tag)
		}
	}

	// Construct the layout result.
	result := results.LayoutResult{
		Page:    "", // Can be set by the caller if needed.
		Status:  "pass",
		Message: "Layout is consistent.",
	}
	if len(missingElements) > 0 {
		result.Status = "fail"
		result.Message = "Missing expected elements: " + strings.Join(missingElements, ", ")
	}

	return result
}
