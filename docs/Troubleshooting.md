# Troubleshooting Guide

This guide helps you diagnose and resolve common issues when using CedarHawk.

## Configuration Issues

- **Error:** "Failed to load config" or "required field 'startURL' is missing"
  - **Solution:**  
    Ensure your configuration file is correctly formatted in TOML and includes all required fields (e.g., `startURL`). Validate the file using an online TOML validator if needed.

## Network and Crawler Issues

- **Error:** Crawler fails to fetch pages or returns non-OK HTTP status codes.
  - **Solution:**  
    Check your network connection and verify that the target website is accessible. Look at the error logs in the JSON output for details.

## UI Test Failures

- **Issue:** Layout discrepancies or missing UI elements are reported.
  - **Solution:**  
    Compare your website’s HTML structure with the expected layout defined in your configuration. Adjust the configuration if your site's layout has custom elements.

## JSON Report Issues

- **Issue:** No JSON report is generated or it is incomplete.
  - **Solution:**  
    Confirm that the `resultsFile` parameter in your configuration is correctly set and that the directory has proper write permissions.

## Build and Installation Issues

- **Error:** Compilation errors or missing dependencies.
  - **Solution:**  
    Ensure you are using a supported version of Go and that your environment is properly set up. Run `go mod tidy` to resolve dependencies.

## Logging and Error Capture

- **Issue:** Errors are not being logged as expected.
  - **Solution:**  
    Verify that the logging level is set appropriately in your configuration and that the centralized logging system is initialized in `cmd/main.go`.

If you continue to experience issues, please refer to the FAQ or contact support.

