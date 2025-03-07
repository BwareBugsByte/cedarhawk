package results

import (
	"encoding/json"
	"os"
)

// TestResults represents the aggregated test results.
type TestResults struct {
	Page                  string                 `json:"page"`
	LayoutResults         []LayoutResult         `json:"layoutResults,omitempty"`
	UIElementsResults     []UIElementResult      `json:"uiElementsResults,omitempty"`
	ResponsivenessResults []ResponsivenessResult `json:"responsivenessResults,omitempty"`
	ErrorLogs             []ErrorLog             `json:"errorLogs,omitempty"`
	AIAnalysis            *AIAnalysisReport      `json:"aiAnalysis,omitempty"`
}

// LayoutResult represents the outcome of layout tests.
type LayoutResult struct {
	Page    string `json:"page"`
	Status  string `json:"status"`  // e.g., "pass" or "fail"
	Message string `json:"message"` // description of the result
}

// UIElementResult represents the validation results of a UI element.
type UIElementResult struct {
	ElementType       string `json:"elementType"`                 // e.g., "button", "link", "image", "textField"
	Status            string `json:"status"`                      // "pass" or "fail"
	Message           string `json:"message"`                     // error or success message
	ElementIdentifier string `json:"elementIdentifier,omitempty"` // e.g., an id or CSS selector
}

// ResponsivenessResult represents the results of testing a particular viewport.
type ResponsivenessResult struct {
	Viewport string   `json:"viewport"` // e.g., "desktop", "tablet", "mobile"
	Status   string   `json:"status"`   // "pass" or "fail"
	Issues   []string `json:"issues,omitempty"`
}

// ErrorLog captures an error encountered during testing.
type ErrorLog struct {
	Module    string `json:"module"`    // Name of the module (e.g., "Crawler")
	Message   string `json:"message"`   // Error details
	Timestamp string `json:"timestamp"` // ISO8601 timestamp
}

// AIAnalysisReport holds the AI plugin's analysis.
type AIAnalysisReport struct {
	Summary         string   `json:"summary"`
	Recommendations []string `json:"recommendations,omitempty"`
}

// WriteJSON serializes the TestResults structure to a JSON file at the given filePath.
func WriteJSON(results TestResults, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(results)
}
