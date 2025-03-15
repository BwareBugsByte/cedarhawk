package logging

import (
	"log"
	"os"
	"sync"
	"time"

	"cedarhawk/internal/results"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
	CRITICAL
)

var currentLogLevel = INFO
var mu sync.Mutex

// capturedErrors stores errors captured by logging functions.
var capturedErrors []results.ErrorLog
var errorsMu sync.Mutex

// SetLogLevel sets the current log level.
func SetLogLevel(level LogLevel) {
	mu.Lock()
	defer mu.Unlock()
	currentLogLevel = level
}

// shouldLog returns true if messages with the given level should be logged.
func shouldLog(level LogLevel) bool {
	mu.Lock()
	defer mu.Unlock()
	return level >= currentLogLevel
}

// Debug logs a debug message.
func Debug(module, message string) {
	if shouldLog(DEBUG) {
		log.Printf("[DEBUG] [%s] %s", module, message)
	}
}

// Info logs an info message.
func Info(module, message string) {
	if shouldLog(INFO) {
		log.Printf("[INFO] [%s] %s", module, message)
	}
}

// Warning logs a warning message.
func Warning(module, message string) {
	if shouldLog(WARNING) {
		log.Printf("[WARNING] [%s] %s", module, message)
	}
}

// Error logs an error message and captures it.
func Error(module, message string) {
	if shouldLog(ERROR) {
		log.Printf("[ERROR] [%s] %s", module, message)
	}
	CaptureError(module, message)
}

// Critical logs a critical message, captures it, and may trigger a termination.
func Critical(module, message string) {
	if shouldLog(CRITICAL) {
		log.Printf("[CRITICAL] [%s] %s", module, message)
	}
	CaptureError(module, message)
	// Optionally, you might exit the program here.
}

// CaptureError appends an error log to the global capturedErrors slice.
func CaptureError(module, message string) {
	errorsMu.Lock()
	defer errorsMu.Unlock()
	errLog := results.ErrorLog{
		Module:    module,
		Message:   message,
		Timestamp: time.Now().Format(time.RFC3339),
	}
	capturedErrors = append(capturedErrors, errLog)
}

// GetCapturedErrors returns the slice of captured error logs.
func GetCapturedErrors() []results.ErrorLog {
	errorsMu.Lock()
	defer errorsMu.Unlock()
	return capturedErrors
}

// ResetCapturedErrors clears the global capturedErrors slice.
// This is exported for testing purposes.
func ResetCapturedErrors() {
	errorsMu.Lock()
	defer errorsMu.Unlock()
	capturedErrors = []results.ErrorLog{}
}

// InitializeLogger sets up the logger with standard flags and outputs to stdout.
func InitializeLogger() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags) // Include date and time in logs.
}
