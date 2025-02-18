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
