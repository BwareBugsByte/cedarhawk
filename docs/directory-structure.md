Design: Project Directory Structure and File Organization
================================================
# 1. Directory Structure

The CedarHawk repository is organized into several key directories, each with a well-defined purpose. Below is the proposed structure:

    CedarHawk/
    ├── cmd/
    │   └── main.go                  # Main entry point for the application
    ├── internal/
    │   ├── config/
    │   │   └── config.go            # Configuration loader and parser (TOML)
    │   ├── crawler/
    │   │   └── crawler.go           # Recursive website crawler module
    │   ├── ui/
    │   │   ├── layout.go            # UI testing: Layout Verification module
    │   │   ├── elements.go          # UI testing: UI Elements Validation module
    │   │   └── responsiveness.go    # UI testing: Responsiveness Testing module
    │   ├── results/
    │   │   └── results.go           # JSON results aggregation and generation
    │   ├── logging/
    │   │   └── logging.go           # Unified logging and error handling
    │   └── ai/
    │       └── ai.go                # Placeholder for future AI Plugin integration
    ├── configs/
    │   └── config.toml              # Sample user configuration file
    ├── tests/
    │   └── <various test files>     # Unit and integration tests
    └── docs/
        ├── architecture.md          # System architecture overview documentation
        └── directory-structure.md   # Detailed project directory and naming conventions documentation

Explanation:

    cmd/: Contains the application's main entry point (main.go). This is where the command-line interface (CLI) is implemented, which orchestrates all other modules.
    internal/: Houses all core modules of CedarHawk. Each sub-package (e.g., config, crawler, ui, results, logging, ai) is designed to be self-contained, promoting modularity and ease of testing.
        config/: Responsible for reading and parsing the config.toml file.
        crawler/: Implements the recursive crawling of the target website.
        ui/: Contains sub-modules for various UI tests (layout, element validation, responsiveness).
        results/: Aggregates outputs from tests into a JSON file.
        logging/: Provides a centralized logging mechanism and error handling.
        ai/: Serves as a placeholder for future AI integration.
    configs/: Stores configuration files. The sample config.toml here will allow users to customize their testing parameters.
    tests/: Contains all tests (unit and integration) ensuring that each module functions as expected.
    docs/: Contains design documents, including the system architecture overview and detailed directory structure, to facilitate understanding and future maintenance.

---

# 2. Naming Conventions
Files and Directories:

    File Names:
        Use lowercase letters with descriptive names (e.g., crawler.go, layout.go).
        Avoid spaces; use hyphens or underscores if necessary.
    Directories:
        Name directories based on their function (e.g., cmd, internal, configs, tests, docs).

Variables and Functions:

    Variables:
        Use camelCase for variable names.
        Use descriptive names that clearly indicate the variable’s purpose.
    Functions:
        Use camelCase for function names.
        Public functions (exported) should start with an uppercase letter, while internal helper functions should start with a lowercase letter.
    Types:
        Use PascalCase for type names and struct definitions.

General Guidelines:

    Modularity:
        Each module is self-contained, reducing interdependencies and making the codebase easier to maintain and extend.
    Readability:
        Files and functions should be self-explanatory. Include comments and documentation where necessary.
    Consistency:
        Maintain a consistent naming and directory structure across the project. This helps new developers quickly understand the codebase.

---

# 3. Benefits of This Organization

    Modularity:
        Separating core functionalities into dedicated packages (config, crawler, ui, etc.) allows for independent development and testing of each component.
    Ease of Maintenance:
        Clear organization facilitates troubleshooting and future enhancements. Each module can be modified without impacting others.
    Scalability:
        A well-defined structure makes it easier to integrate additional features (like the AI plugin) as the project evolves.
    Separation of Concerns:
        Different responsibilities are clearly segregated. For example, configuration handling is isolated in its own package, which means changes to configuration parsing won’t affect the crawler or UI testing logic.
