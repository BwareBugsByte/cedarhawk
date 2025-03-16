# Configuration Guide

CedarHawk uses a TOML configuration file to define testing parameters. This document explains the structure of the configuration file and the available options.

## Default Location
The default configuration file is located at:

configs/config.toml


## Example Configuration File

```toml
# Base URL for crawling
startURL = "https://example.com"

# Maximum depth for recursive crawling
maxDepth = 3

# File to output the JSON results
resultsFile = "results.json"

[ui]
# Expected layout identifier (e.g., "default")
expectedLayout = "default"
# Define viewports for responsiveness testing
viewports = [
    { name = "desktop", width = 1920, height = 1080 },
    { name = "tablet", width = 768, height = 1024 },
    { name = "mobile", width = 375, height = 667 }
]

[logging]
# Logging level (e.g., "info", "debug")
level = "info"
```

## Field Descriptions

  - startURL:
    The base URL from which the crawler will begin its tests.

  - maxDepth:
    The maximum recursion depth for the crawler to avoid infinite loops.

  - resultsFile:
    The output file for the consolidated JSON test report.

    [ui] Section:
        expectedLayout:
        The layout pattern that the UI tests expect (e.g., "default").
        viewports:
        A list of viewport configurations for responsiveness testing. Each viewport has:
            name: A label (e.g., desktop, mobile).
            width, height: Dimensions to simulate.

    [logging] Section:
        level:
        The desired logging level (e.g., "info", "debug").

## Customization

To use a custom configuration file, supply the --config flag when running CedarHawk:
```bash
./cedarhawk --config /path/to/your/config.toml
```
