package config

// Config represents the parsed configuration.
type Config struct {
	StartURL    string
	MaxDepth    int
	ResultsFile string
	UI          UIConfig
	Logging     LoggingConfig
}

// UIConfig represents UI-specific configuration.
type UIConfig struct {
	ExpectedLayout string
	Viewports      []Viewport
}

// Viewport represents a screen size for responsiveness testing.
type Viewport struct {
	Name   string
	Width  int
	Height int
}

// LoggingConfig represents logging configuration.
type LoggingConfig struct {
	Level string
}

// LoadConfig is a placeholder for the config loader.
func LoadConfig(path string) (*Config, error) {
	// TODO: Implement TOML parsing using built-in packages.
	return &Config{
		StartURL:    "https://example.com",
		MaxDepth:    3,
		ResultsFile: "results.json",
		UI: UIConfig{
			ExpectedLayout: "default",
			Viewports: []Viewport{
				{"desktop", 1920, 1080},
				{"tablet", 768, 1024},
				{"mobile", 375, 667},
			},
		},
		Logging: LoggingConfig{
			Level: "info",
		},
	}, nil
}
