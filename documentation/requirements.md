CedarHawk Requirements:
=======================
***Project Requirements, Goals, and Scope***

**Author** *BwareBugsByte*

# Project Setup and Configuration
Language and Libraries

    GoLang: The project must be entirely programmed in GoLang due to its simplicity, speed, and scalability.

    Built-in Libraries: Minimize the use of external libraries and tools. Only use built-in GoLang libraries to reduce the risk of dependency issues.

Project Goal

    Ease of Use: The program should function seamlessly when a user clones the repo and has Go installed.

File and Directory Structure

    Naming Conventions: All file names should be self-explanatory.

    Organization: Use directories to categorize files. Each "thing" (file, variable, function) will have:

        Category: Describes the group a "thing" belongs to (i.e., the folder or directory).

        Tag: Descriptive tokens derived from the title. Ensure tags are limited for sanity.

Configuration Files

    Input Configuration: Use TOML format for user-generated configuration files due to its readability.

    Output Configuration: Use JSON format for outputted configuration files or data for easy parsing by computers.
---
#Implementation Details
Description

    Initialize the Go module for CedarHawk.

    Create a clean project structure (e.g., /cmd, /internal, /configs, /tests).

    Set up configuration file support using the TOML format.

        Implement a minimal TOML parser using built-in Go packages, or limit configuration to simple key-value pairs if necessary.

    Implement functionality to generate and write test results in JSON format.

        Use the standard encoding/json package.

        Ensure that results include detailed test information (e.g., passed/failed status, error messages, timestamps).

Acceptance Criteria

    A new Git repository with an initialized Go module.

    A sample config.toml file with documented fields.

    Basic code to read and parse the config file using only built-in tools.

    A module that collects test data and writes a well-formatted JSON file.

    Example JSON output provided in the documentation.
