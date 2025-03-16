Design: Configuration File Handling
======
# 1. Overview

The Configuration File Handling module is responsible for reading and parsing the user-generated configuration file written in TOML format.

---

# 2. Configuration Schema

The configuration file (e.g., config.toml) will include several sections with key-value pairs. A sample schema is outlined below:

# CedarHawk configuration file example

# Global settings
startURL = "https://example.com"  # (Required) Base URL for the crawler
maxDepth = 3                      # (Optional) Maximum recursion depth for crawling; default: 3
resultsFile = "results.json"      # (Optional) Path for the JSON output file

[ui]
expectedLayout = "default"        # (Optional) Identifier for the expected layout, can be extended
viewports = [
    { name = "desktop", width = 1920, height = 1080 },
    { name = "tablet", width = 768, height = 1024 },
    { name = "mobile", width = 375, height = 667 }
]

[logging]
level = "info"                    # (Optional) Logging level (e.g., "info", "debug"); default: "info"

Required Fields

    Global:
        startURL: The starting URL for crawling (required).

Optional Fields with Defaults

    Global:
        maxDepth: Default is 3 if not provided.
        resultsFile: Default file name for output.
    UI Section:
        expectedLayout: Can default to a predefined value like "default".
        viewports: A list of viewports for responsiveness testing.
    Logging Section:
        level: Default logging level set to "info".

---

3. Parsing Strategy
Step-by-Step Approach

    File Reading:
        Open the configuration file using os.Open and create a bufio.Scanner to process it line by line.

    Line Processing:
        Ignore Comments & Empty Lines:
            Skip lines that are empty or start with #.
        Section Detection:
            Detect section headers by checking for lines that start with [ and end with ].
            Maintain a current section context to prefix keys accordingly.
        Key-Value Parsing:
            Split lines containing a key-value pair using the first occurrence of =.
            Trim whitespace around keys and values.
        Array Handling:
            For keys that start with [ (like the viewports array), implement minimal logic to extract individual table elements.
            Support only the required subset (i.e., array of inline tables with simple key-value pairs).

    Mapping to Go Structures:
        Populate a predefined Go struct (or nested structs) based on the configuration schema.
        Convert string values to the appropriate types (e.g., integers for maxDepth, width, and height).

---

# 4. Error Handling and Default Values

    Missing Required Fields:
        If a required field (such as startURL) is missing, the parser will return an error and halt further processing.
    Default Value Assignment:
        For optional fields that are not provided, default values will be assigned (e.g., maxDepth defaults to 3).
    Malformed Lines:
        If a line does not conform to the expected format (e.g., missing =), log an error using the Logging module and either skip the line or apply a default, based on severity.
    Section-Specific Errors:
        If a section (e.g., [ui] or [logging]) is malformed or missing entirely, the parser should apply the corresponding default configuration for that section.

---

# 5. Data Flow Diagram

              +------------------------+
              |    User config.toml    |
              +-----------+------------+
                          |
                          v
              +-----------+-----------+
              | Configuration Loader  | 
              +-----------+-----------+
                          |
                          v
              +-----------+-----------+
              |  Config Struct (Go)   |
              +-----------+-----------+
                          |
                          v
              +-----------+-----------+
              |    Other Modules      |
              | (UI Testing, Crawler, |
              |    Logging, etc.)     |
              +-----------------------+
