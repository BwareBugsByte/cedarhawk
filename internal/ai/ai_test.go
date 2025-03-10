package ai

import (
	"testing"

	"cedarhawk/internal/results"
)

func TestDefaultAIPluginStub(t *testing.T) {
	// Create a stub instance.
	stub := &DefaultAIPluginStub{}

	// Create a dummy TestResults struct.
	testResults := &results.TestResults{
		Page: "https://example.com",
	}

	// Invoke the AnalyzeResults method.
	report, err := stub.AnalyzeResults(testResults)
	if err != nil {
		t.Fatalf("Expected no error from AnalyzeResults, got: %v", err)
	}

	// Verify the returned analysis report has the expected placeholder message.
	expectedSummary := "AI analysis not available."
	if report.Summary != expectedSummary {
		t.Errorf("Expected report summary '%s', got '%s'", expectedSummary, report.Summary)
	}

	// Verify that recommendations are empty.
	if report.Recommendations == nil || len(report.Recommendations) != 0 {
		t.Errorf("Expected no recommendations, got: %v", report.Recommendations)
	}
}

