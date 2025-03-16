Design: Main Entry Point and Module Integration
============================
# 1. Overview

The Main Entry Point serves as the centralized command-line interface (CLI) for CedarHawk. It is responsible for coordinating the execution of all modules—from loading configuration and crawling the website to running UI tests and aggregating results into a JSON file. The design ensures that each module can be executed either independently or as part of the full test suite, and that errors are propagated and reported in a controlled manner.

---

# 2. Main CLI Design
Location and Structure

    File Location:
        The main entry point will reside in cmd/main.go.
    Command-Line Flags:
        Flags will allow users to select specific modules or run the full suite. For example:
            --config to specify a configuration file.
            --crawl-only to run only the crawler.
            --ui-only to run only UI tests.
            --all (default) to run all modules sequentially.

High-Level Responsibilities

    Configuration Loading:
        Parse and load user configuration using the Configuration Loader.
    Module Coordination:
        Sequentially (or optionally in parallel) call the Crawler Module, UI Testing Modules, and Results Generator.
    Error Handling:
        Capture and log errors from each module.
        Allow non-critical errors to be included in the final report without aborting the entire process.
    Result Aggregation:
        Collect outputs from all modules into a unified TestResults structure.
    Optional AI Integration:
        Invoke the AI Plugin Integration Interface after aggregating results.
    JSON Reporting:
        Write the final results to a JSON file as specified by the configuration.

---

# 3. Sequence of Operations
Step-by-Step Flow:

    Startup and Flag Parsing:
        The main function parses command-line flags.
        If a configuration file is not specified, a default path is used.

    Configuration Loading:
        Load and validate the configuration file (TOML) using the Configuration Loader.
        Exit with an error if a critical configuration (e.g., startURL) is missing.

    Module Execution:
        Crawler Module:
            Start crawling from the base URL.
            Collect a list of unique URLs and associated HTML/DOM data.
            If the --crawl-only flag is provided, output the crawler results and terminate.
        UI Testing Modules:
            Layout Verification: Compare page structures across URLs.
            UI Elements Validation: Check for the presence and functionality of essential UI elements.
            Responsiveness Testing: Simulate various viewport sizes to verify layout adaptation.
        Results Aggregation:
            Combine the outputs from the crawler and UI testing modules into a single TestResults structure.
        AI Plugin (Optional):
            If an AI plugin is configured, invoke its AnalyzeResults method.
            If not, the default stub logs “AI plugin not implemented.”

    JSON Result Generation:
        Serialize the TestResults structure to JSON.
        Write the JSON to the file specified in the configuration (e.g., results.json).

    Exit and Reporting:
        Log a summary of the execution.
        Exit with a status code that reflects success or failure based on critical errors.

---

# 4. Error Propagation and Reporting

    Centralized Logging:
        A dedicated Logging Module will capture errors and warnings during execution.
    Module-Level Error Handling:
        Each module will return detailed error information along with its results.
        The main function will aggregate these error messages and include them in the JSON report.
    Non-Critical Failures:
        Modules are designed to continue processing where possible, logging issues rather than aborting execution.
    Critical Failures:
        If a critical module (e.g., Configuration Loader) fails, the CLI will exit immediately with an error code.

---

# 5. Independent and Combined Module Execution
Independent Execution:

    Flags:
        Command-line flags (e.g., --crawl-only or --ui-only) allow users to run specific modules.
    Module Isolation:
        Each module exposes its own public function that can be invoked separately. For instance:
            RunCrawler(config) returns crawl results.
            RunUITests(htmlData, config) returns UI test results.

Combined Execution:

    Default Behavior:
        Without specific flags, the main entry point runs all modules in sequence:
            Load configuration.
            Execute crawler.
            Run UI tests on the crawled pages.
            Aggregate and generate JSON results.
    Extensibility:
        The modular design facilitates adding new modules (e.g., future AI enhancements) with minimal changes to the main logic.

---

# 6. Integration Diagram

               +-------------------------+
               |      Main Entry         |
               |       (cmd/main.go)     |
               +------------+------------+
                            |
                +-----------v-----------+
                | Configuration Loader  |
                +-----------+-----------+
                            |
                            v
                +-----------+-----------+
                |  Crawler Module       |
                +-----------+-----------+
                            |
                            v
            +---------------+---------------+
            |         UI Testing Modules    |
            |  (Layout, UI Elements,        |
            |   Responsiveness Testing)     |
            +---------------+---------------+
                            |
                            v
                +-----------+-----------+
                | Results Aggregator    |
                +-----------+-----------+
                            |
                            v
                +-----------+-----------+
                | AI Plugin Integration | (Optional)
                +-----------+-----------+
                            |
                            v
                +-----------+-----------+
                | JSON Report Generator |
                +-----------------------+

---

# 7. Sample Pseudocode

func main() {
    // Parse command-line flags (config file path, module selection, etc.)
    configPath := parseFlags()

    // Load configuration
    config, err := config.LoadConfig(configPath)
    if err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }

    // Execute crawler module
    crawlResults, err := crawler.Run(config.StartURL, config.MaxDepth)
    if err != nil {
        log.Printf("Crawler error: %v", err)
    }

    // Optionally run only the crawler
    if flagCrawlOnly {
        outputJSON(crawlResults)
        return
    }

    // Run UI Testing Modules using HTML data from crawler
    layoutResults := ui.RunLayoutVerification(crawlResults.HTMLData, config.UI)
    uiElementsResults := ui.RunElementsValidation(crawlResults.HTMLData, config.UI)
    responsivenessResults := ui.RunResponsivenessTests(crawlResults.HTMLData, config.UI.Viewports)

    // Aggregate all results into a unified structure
    testResults := aggregateResults(crawlResults, layoutResults, uiElementsResults, responsivenessResults)

    // Optionally invoke AI plugin analysis
    aiReport, err := aiPlugin.AnalyzeResults(testResults)
    if err != nil {
        log.Printf("AI analysis error: %v", err)
    } else {
        testResults.AIAnalysis = aiReport
    }

    // Generate JSON output
    if err := results.WriteJSON(testResults, config.ResultsFile); err != nil {
        log.Printf("Failed to write JSON results: %v", err)
    }

    // Final summary logging and exit
    log.Println("CedarHawk tests completed successfully.")
}

