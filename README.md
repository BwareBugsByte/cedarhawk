# CedarHawk

CedarHawk is an automated website functionality testing tool written in Go. It performs recursive crawling, UI testing (layout, elements, responsiveness), and optionally integrates an AI analysis module to verify website functionality.

## Features
- **Recursive Crawler:** Automatically traverses the website and gathers HTML.
- **UI Testing:** Verifies layout consistency, essential UI elements, and responsiveness.
- **AI Plugin Stub:** Provides a placeholder for future AI-based validation.
- **JSON Reporting:** Aggregates test results and errors into a consolidated JSON file.

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/BwareBugsByte/cedarhawk.git
   cd cedarhawk
   ```

2. Build the binary:
   ```bash
    go build ./cmd/main.go
   ```
## Usage
Run CedarHawk with default settings:
```bash
./cedarhawk
```
For help:
```bash
./cedarhawk --help
```
Command-line flags:

  - --config: Specify a configuration file (default: configs/config.toml)
  - --crawl-only: Run only the crawler module.
  - --ui-only: Run only the UI tests.

## Configuration

Edit the configs/config.toml file to set parameters such as:

  - startURL: The base URL for crawling.
  - maxDepth: Maximum depth for crawling.
  - UI settings: Expected layout and viewports.
  - Logging levels.

## Documentation

Additional guides and troubleshooting instructions can be found in the /docs folder.
## Troubleshooting

  - Ensure Go is installed and your GOPATH is set.
  - Verify that the configuration file is correctly formatted.
  - Consult the /docs folder for FAQs and troubleshooting tips.
