CedarHawk Project Requirements
=======================

***Author:*** **BwareBugsByte**

# 1. Project Setup and Configuration
Language and Libraries

    GoLang: Entire project must be programmed in GoLang for simplicity, speed, and scalability.
    Built-in Libraries: Minimize the use of external libraries. Prefer built-in GoLang libraries to reduce dependency issues.

Project Goal

    Ease of Use: The program should function seamlessly when a user clones the repository and has Go installed.

File and Directory Structure

    Naming Conventions: All file names must be self-explanatory.
    Organization: Use directories to categorize files. Each item (file, variable, function) will have:
        Category: Describes the group an item belongs to (i.e., folder or directory).
        Tag: Descriptive tokens derived from the title with limited tags for sanity.

Configuration Files

    Input Configuration: Use TOML format for user-generated configuration files due to its readability.
    Output Configuration: Use JSON format for outputted configuration files or data for easy parsing by computers.

---

# 2. Implementation Details
Description

    Initialize Go Module: Set up the Go module for CedarHawk.
    Project Structure: Create a clean structure (e.g., /cmd, /internal, /configs, /tests).

Configuration File Support

    Implement support for TOML format.
    Use built-in Go packages for minimal TOML parsing or limit to simple key-value pairs.

Test Results in JSON

    Generate and write test results using encoding/json.
    Include detailed test information (passed/failed status, error messages, timestamps).

Crawler

    Start at a given URL and recursively fetch and parse HTML pages.
    Use net/http and built-in libraries to process HTML.
    Avoid loops and duplicate URLs.

Layout Verification

    Verify layout consistency across different pages and screen sizes.
    Extract DOM structure (e.g., container hierarchy, common elements) and compare across pages.
    Configure expected layout patterns via config.toml.

Responsiveness Testing

    Simulate various viewport sizes (desktop, tablet, mobile).
    Ensure elements scale appropriately and layout adapts without breaking.

UI Elements Validation

    Validate the presence of essential UI elements (buttons, links, images, text fields) on each page.
    Verify image load success (detect broken links).
    Optionally check for functional attributes (e.g., onclick handlers on buttons).

AI Plugin Integration Placeholder

    Design the system to allow seamless integration of an AI module at a later stage.
    Stub functionality to log “AI plugin not implemented” without affecting other tests.

Main Entry Point

    Automatically load configurations, initiate tests (crawling, UI testing, etc.), and aggregate results.
    Ensure a modular design for independent or combined execution of testing components.

Acceptance Criteria

    Go Module: Initialized in a new Git repository.
    Configuration: A sample config.toml file with documented fields.
    Config Parsing: Basic code to read and parse the config file using built-in tools.
    JSON Results: Module to collect test data and generate a well-formatted JSON file, with example output provided in the documentation.
    Crawler: A functional crawler that visits and stores a list of unique URLs, with the ability to extract HTML DOM elements for further testing.
    Layout Verification: A tool comparing layout structures and flagging inconsistencies, with a JSON report detailing any differences.
    Responsiveness Testing: A tool that runs UI tests across different viewport dimensions, with a JSON report indicating responsiveness issues.
    UI Elements Validation: A module logging missing or malfunctioning UI components, with a JSON report summarizing the validations.
    AI Plugin Placeholder: A documented interface/placeholder for AI plugin integration, with stub functionality logging “AI plugin not implemented.”
    Main Execution: Single command execution performing all tests and writing output to a JSON file.
    Modularity: Documented modular design with instructions to add or remove modules.
