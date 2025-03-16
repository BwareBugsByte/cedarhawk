# Usage Guide

This document provides instructions on how to run CedarHawk and details the available command-line options.

## Running CedarHawk

After building the binary, you can run CedarHawk from the command line:
```bash
./cedarhawk
```

## Command-Line Flags

  - --config
    Specify the path to a configuration file.
      - Example:
      ```bash
      ./cedarhawk --config configs/custom_config.toml
      ```

  - --crawl-only
    Run only the crawler module.
      - Example:
      ```bash
      ./cedarhawk --crawl-only
      ```

  - --ui-only
    Run only the UI testing modules.
    - Example:
    ```bash
    ./cedarhawk --ui-only
    ```

  - --help
    Display usage instructions and available flags.
    - Example:
    ```bash
    ./cedarhawk --help
    ```

## Output

CedarHawk generates a consolidated JSON report (by default results.json) that contains:

  - Crawled page data
  - UI test results (layout verification, UI elements validation, responsiveness tests)
  - Error logs and optional AI analysis report

Review the JSON output to analyze the test results.
Example

To run the full test suite:
```bash
./cedarhawk
```

To run only the crawler:
```bash
./cedarhawk --crawl-only
```

To view help:
```bash
./cedarhawk --help
```
