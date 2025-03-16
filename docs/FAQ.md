FAQ.md

# Frequently Asked Questions (FAQ)

This document addresses common questions about CedarHawk.

## Q1: What is CedarHawk?
**A:** CedarHawk is an automated website functionality testing tool written in Go. It performs recursive crawling, UI testing (layout verification, UI elements validation, responsiveness testing), and generates a consolidated JSON report.

## Q2: How do I install CedarHawk?
**A:** Clone the repository and build the binary using the instructions in [Installation.md](Installation.md).

## Q3: How do I configure CedarHawk?
**A:** CedarHawk uses a TOML configuration file located in the `/configs` folder by default. See [Configuration.md](Configuration.md) for details on configuration options.

## Q4: What command-line flags are available?
**A:** You can use flags such as `--config`, `--crawl-only`, `--ui-only`, and `--help`. For a complete list, run:
```bash
./cedarhawk --help
```

## Q5: Where is the JSON report generated?
**A:** The JSON report is created at the location specified by the resultsFile parameter in your configuration (default: results.json).

## Q6: How do I troubleshoot issues with CedarHawk?
**A:** Refer to the Troubleshooting.md guide for solutions to common problems.

## Q7: Can I contribute to CedarHawk?
**A:** Yes, contributions are welcome! Please see the repository's contribution guidelines for more information.

## Q8: Who do I contact for support?
**A:** For support, open an issue on the GitHub repository
