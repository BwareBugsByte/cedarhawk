Design: Logging, Error Handling, and Reporting
==========================
# 1. Overview

This document outlines a unified strategy for logging and error handling across all CedarHawk modules. The design focuses on consistency in log output, clear error propagation, and integration of error details into the final JSON results. It caters to both development and production environments by supporting configurable log levels and reporting formats.

---

# 2. Unified Logging Strategy
Log Levels and Format

    Log Levels:
        DEBUG: Detailed information, typically of interest only when diagnosing problems.
        INFO: General operational messages that highlight the progress of the application.
        WARNING: Indications of possible issues that do not halt the system.
        ERROR: Significant issues that affect functionality and require attention.
        CRITICAL: Severe errors that may cause the program to abort.
    Format:
        Logs will include timestamps, log level, source module, and the message.
        Example: 2025-02-18T10:23:45Z [INFO] [Crawler] Successfully crawled URL: https://example.com

Implementation Using Built-In Packages

    Go’s log Package:
        The design leverages the built-in log package.
        Custom wrapper functions can be created to standardize output and manage log levels.
    Configuration-Driven Logging:
        Log level is specified in the config.toml (e.g., level = "info" in the [logging] section).
        In development, lower log levels (DEBUG, INFO) can be enabled, while production may limit to WARNING and above.

---

# 3. Error Handling Strategy
Module-Level Error Reporting

    Error Propagation:
        Each module (e.g., Configuration Loader, Crawler, UI Testing Modules) is responsible for returning errors encountered during execution.
        Errors are propagated to the main entry point, where they are aggregated.
    Granular Error Information:
        Errors should include context such as the module name, function, and a descriptive message.
        For example: "Crawler error: failed to fetch URL https://example.com (HTTP 404)"

Error Capture in JSON Results

    Results Aggregation:
        The final TestResults structure includes an errors field.
        Each error is stored as an object with details such as module name, error message, and timestamp.
    Sample JSON Structure:

    {
      "page": "https://example.com",
      "results": { /* test results */ },
      "errors": [
        {
          "module": "Crawler",
          "message": "Failed to fetch URL https://example.com (HTTP 404)",
          "timestamp": "2025-02-18T10:23:45Z"
        },
        {
          "module": "UI Elements Validation",
          "message": "Missing onclick attribute for button btn-submit",
          "timestamp": "2025-02-18T10:24:10Z"
        }
      ]
    }

---

# 4. Unified Logging Implementation
Logging Interface and Wrapper Functions

    Wrapper Functions:
        Create functions such as LogDebug(module, message), LogInfo(module, message), etc., to standardize logging output.
        These functions will check the current log level and output the log message accordingly.
    Centralized Logging Module:
        All modules import the centralized logging module (e.g., internal/logging/logging.go).
        This ensures that log messages across modules follow the same format and settings.

Error Handling Integration

    Standard Error Handling Pattern:
        Each module returns an error along with its result.
        The main entry point aggregates these errors.
    Non-Critical vs. Critical Errors:
        Non-critical errors are logged and recorded in the JSON results.
        Critical errors (e.g., missing configuration) trigger immediate termination, with an appropriate log and error message.

---

# 5. Environment Strategies
Development Environment:

    Verbose Logging:
        Log levels set to DEBUG or INFO.
        Detailed logs aid in diagnosing issues during development.
    Error Reporting:
        Errors are logged with full stack/contextual information.

Production Environment:

    Controlled Logging:
        Log levels set to WARNING or ERROR to reduce noise.
        Logs can be redirected to a file or external logging system.
    Resilience:
        The system is designed to continue operation despite non-critical errors.
        Critical errors trigger a safe shutdown with detailed log reporting.

---

# 6. Integration with JSON Reporting

    Error Aggregation:
        The Results Generator module collects errors from all modules and integrates them into the final JSON output.
    Unified Output:
        Both successful test results and error logs are presented in one consolidated JSON file.
    Logging Integration:
        As each module logs messages, these logs are stored in a common log file (or console) for immediate review, while the error summary in the JSON results provides a permanent record for later analysis.

---

# 7. Diagram: Logging and Error Handling Flow

                +------------------------+
                |  Module Operations     |
                |  (Crawler, UI Tests,    |
                |   Responsiveness, etc.)|
                +-----------+------------+
                            |
             +--------------v---------------+
             |  Centralized Logging Module  |
             | (Log levels, Wrapper funcs)  |
             +--------------+---------------+
                            |
                            v
             +--------------+---------------+
             |  Error Handling in Modules   |
             |  (Return errors with context)|
             +--------------+---------------+
                            |
                            v
             +--------------+---------------+
             |  Main Entry Point Aggregates |
             |     Errors & Test Results    |
             +--------------+---------------+
                            |
                            v
             +--------------+---------------+
             |   JSON Results Generator     |
             | (Integrates error summaries) |
             +------------------------------+

