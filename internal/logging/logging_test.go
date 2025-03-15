package logging

import (
	"strings"
	"testing"

	"cedarhawk/internal/results"
)

// Force usage of the results package.
var _ = results.ErrorLog{}

func TestCaptureError(t *testing.T) {
	ResetCapturedErrors()

	// Call Error to capture an error.
	Error("TestModule", "This is a test error")
	captured := GetCapturedErrors()
	if len(captured) != 1 {
		t.Errorf("Expected 1 captured error, got %d", len(captured))
	}
	if captured[0].Module != "TestModule" {
		t.Errorf("Expected module 'TestModule', got '%s'", captured[0].Module)
	}
	if !strings.Contains(captured[0].Message, "This is a test error") {
		t.Errorf("Expected error message to contain 'This is a test error', got '%s'", captured[0].Message)
	}
}

func TestLogLevels(t *testing.T) {
	// Set log level to WARNING so that DEBUG and INFO messages won't be logged.
	SetLogLevel(WARNING)
	ResetCapturedErrors()

	// These calls should not capture errors.
	Debug("TestModule", "Debug message")
	Info("TestModule", "Info message")

	// These should capture errors.
	Warning("TestModule", "Warning message")
	Error("TestModule", "Error message")
	Critical("TestModule", "Critical message")

	captured := GetCapturedErrors()
	// Expect errors from Error and Critical only.
	if len(captured) != 2 {
		t.Errorf("Expected 2 captured errors, got %d", len(captured))
	}
	// Verify captured messages.
	if !strings.Contains(captured[0].Message, "Error message") {
		t.Errorf("Expected first captured error to contain 'Error message', got '%s'", captured[0].Message)
	}
	if !strings.Contains(captured[1].Message, "Critical message") {
		t.Errorf("Expected second captured error to contain 'Critical message', got '%s'", captured[1].Message)
	}
}

