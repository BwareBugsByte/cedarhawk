package ai

import "log"

// AIPlugin defines the interface for an AI module.
type AIPlugin interface {
	AnalyzeResults() string
}

// DefaultAIPluginStub is the stub implementation.
type DefaultAIPluginStub struct{}

// AnalyzeResults logs that the AI plugin is not implemented.
func (d *DefaultAIPluginStub) AnalyzeResults() string {
	log.Println("AI plugin not implemented")
	return "AI analysis not available"
}
