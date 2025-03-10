package ai

import (
	"log"

	"cedarhawk/internal/results"
)

// AIPlugin defines the interface for an AI module.
type AIPlugin interface {
	// AnalyzeResults processes the given test results and returns an AI analysis report.
	AnalyzeResults(results *results.TestResults) (*results.AIAnalysisReport, error)
}

// DefaultAIPluginStub is a stub implementation of the AIPlugin interface.
type DefaultAIPluginStub struct{}

// AnalyzeResults logs that the AI plugin is not implemented and returns a placeholder report.
func (stub *DefaultAIPluginStub) AnalyzeResults(res *results.TestResults) (*results.AIAnalysisReport, error) {
	log.Println("AI plugin not implemented")
	report := &results.AIAnalysisReport{
		Summary:         "AI analysis not available.",
		Recommendations: []string{},
	}
	return report, nil
}


