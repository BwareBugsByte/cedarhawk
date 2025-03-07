package config

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestLoadConfig_Valid(t *testing.T) {
	// Create a temporary file with valid configuration content.
	tmpFile, err := ioutil.TempFile("", "config-*.toml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	content := `
startURL = "https://example.com"
maxDepth = 5
resultsFile = "output.json"

[ui]
expectedLayout = "custom"
[[ui.viewports]]
name = "desktop"
width = 1920
height = 1080

[logging]
level = "debug"
`
	if _, err := tmpFile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	cfg, err := LoadConfig(tmpFile.Name())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if cfg.StartURL != "https://example.com" {
		t.Errorf("expected startURL 'https://example.com', got %s", cfg.StartURL)
	}

	if cfg.MaxDepth != 5 {
		t.Errorf("expected maxDepth 5, got %d", cfg.MaxDepth)
	}

	if cfg.ResultsFile != "output.json" {
		t.Errorf("expected resultsFile 'output.json', got %s", cfg.ResultsFile)
	}

	if cfg.UI.ExpectedLayout != "custom" {
		t.Errorf("expected UI expectedLayout 'custom', got %s", cfg.UI.ExpectedLayout)
	}

	if len(cfg.UI.Viewports) != 1 {
		t.Errorf("expected 1 viewport, got %d", len(cfg.UI.Viewports))
	}

	if cfg.Logging.Level != "debug" {
		t.Errorf("expected logging level 'debug', got %s", cfg.Logging.Level)
	}
}

func TestLoadConfig_MissingRequiredField(t *testing.T) {
	// Create a temporary file without the required startURL.
	tmpFile, err := ioutil.TempFile("", "config-*.toml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	content := `
maxDepth = 5
resultsFile = "output.json"
`
	if _, err := tmpFile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	_, err = LoadConfig(tmpFile.Name())
	if err == nil {
		t.Fatal("expected error for missing startURL, got nil")
	}
}

func TestLoadConfig_MalformedConfig(t *testing.T) {
	// Create a temporary file with malformed configuration.
	tmpFile, err := ioutil.TempFile("", "config-*.toml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	content := `
startURL = "https://example.com"
maxDepth = "not_a_number"
`
	if _, err := tmpFile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	_, err = LoadConfig(tmpFile.Name())
	if err == nil {
		t.Fatal("expected error for malformed config, got nil")
	}
}
