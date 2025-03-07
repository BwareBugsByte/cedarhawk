package config

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml"
)

// Config represents the parsed configuration.
type Config struct {
	StartURL    string        `toml:"startURL"`
	MaxDepth    int           `toml:"maxDepth"`
	ResultsFile string        `toml:"resultsFile"`
	UI          UIConfig      `toml:"ui"`
	Logging     LoggingConfig `toml:"logging"`
}

// UIConfig represents UI-specific configuration.
type UIConfig struct {
	ExpectedLayout string     `toml:"expectedLayout"`
	Viewports      []Viewport `toml:"viewports"`
}

// Viewport represents a screen size for responsiveness testing.
type Viewport struct {
	Name   string `toml:"name"`
	Width  int    `toml:"width"`
	Height int    `toml:"height"`
}

// LoggingConfig represents logging configuration.
type LoggingConfig struct {
	Level string `toml:"level"`
}

// LoadConfig reads and parses the configuration from the given file path.
// It applies default values for optional parameters and validates required fields.
func LoadConfig(path string) (*Config, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("configuration file not found: %s", path)
	}

	// Load and parse the TOML file.
	tree, err := toml.LoadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse configuration file: %v", err)
	}

	// Create a Config instance with default values.
	config := &Config{
		MaxDepth:    3,              // default maxDepth
		ResultsFile: "results.json", // default results file
		UI: UIConfig{
			ExpectedLayout: "default", // default expected layout
		},
		Logging: LoggingConfig{
			Level: "info", // default logging level
		},
	}

	// Unmarshal the TOML tree into the Config struct.
	if err := tree.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal configuration: %v", err)
	}

	// Validate required fields.
	if config.StartURL == "" {
		return nil, fmt.Errorf("required field 'startURL' is missing")
	}

	return config, nil
}
