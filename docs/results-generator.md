Design: Results Generator
=====================
# 1. Overview

The Results Generator module is responsible for aggregating test outputs from all CedarHawk testing components—including the UI Testing Modules, the Crawler Module, and the Logging and Error Handling Module—and then producing a comprehensive JSON report. This JSON file will detail the outcomes of each test, including pass/fail status, error messages, and timestamps.

---

# 2. Role and Responsibilities

    Aggregation:
        Collect outputs from various testing modules (Layout Verification, UI Elements Validation, Responsiveness Testing) and the Crawler.
        Incorporate relevant logging information and error messages captured during the test execution.

    JSON Report Generation:
        Structure the aggregated data into a unified JSON format using Go's built-in encoding/json package.
        Ensure that the JSON report includes essential details such as:
            Test Status: Indicating if each test passed or failed.
            Error Messages: Descriptive messages for any failures or anomalies.
            Timestamps: When each test was executed or when an error occurred.
            Module-Specific Data: Information specific to each module's output.

---

# 3. Data Structures and JSON Schema
Core Data Structure Example (in Go)

type TestResult struct {
    Module      string    `json:"module"`      // Module name (e.g., "Layout Verification", "Crawler")
    Status      string    `json:"status"`      // "pass" or "fail"
    Error       string    `json:"error,omitempty"` // Error message if any; omitted if none
    Timestamp   time.Time `json:"timestamp"`   // Time when the test was executed
    Details     any       `json:"details,omitempty"` // Additional module-specific data (e.g., element counts, discrepancies)
}

type AggregatedResults struct {
    StartTime   time.Time    `json:"start_time"`
    EndTime     time.Time    `json:"end_time"`
    Results     []TestResult `json:"results"`     // List of all test results
    Logs        []string     `json:"logs,omitempty"` // Collected log entries (optional)
}

Example JSON Output

{
  "start_time": "2025-02-18T10:00:00Z",
  "end_time": "2025-02-18T10:05:00Z",
  "results": [
    {
      "module": "Layout Verification",
      "status": "pass",
      "timestamp": "2025-02-18T10:01:00Z",
      "details": {
        "pages_tested": 10,
        "inconsistencies": 0
      }
    },
    {
      "module": "UI Elements Validation",
      "status": "fail",
      "error": "Missing button element on page /contact",
      "timestamp": "2025-02-18T10:02:00Z"
    }
  ],
  "logs": [
    "2025-02-18T10:00:30Z - INFO - Started crawling https://example.com",
    "2025-02-18T10:01:05Z - ERROR - Failed to load image on /about"
  ]
}

---

# 4. Aggregation and Processing Flow
Data Flow Diagram

              +-------------------------+
              |  UI Testing Modules     |
              | (Layout, Elements,      |
              |  Responsiveness)        |
              +------------+------------+
                           |
                           v
              +-------------------------+
              |    Crawler Module       |
              | (HTML & DOM Extraction) |
              +------------+------------+
                           |
                           v
              +-------------------------+
              |  Logging & Error Module |
              +------------+------------+
                           |
                           v
              +-------------------------+
              |   Results Generator     |
              |   (Aggregation & JSON)  |
              +------------+------------+
                           |
                           v
              +-------------------------+
              |     JSON Report File    |
              |    (e.g., results.json) |
              +-------------------------+

Processing Steps

    Collection:
        Each module, upon completion, produces a structured output (e.g., a TestResult instance) containing its findings.
        Log messages and error entries are captured by the Logging Module.

    Aggregation:
        The Results Generator gathers all the TestResult entries and logging data.
        It also records the start and end times of the testing process.

    JSON Marshalling:
        The aggregated results are stored in an AggregatedResults struct.
        The struct is marshalled to JSON using Go's encoding/json package.

    Output:
        The JSON data is written to a file (e.g., results.json) as specified in the configuration.
        In case of errors during the JSON marshalling or file writing process, error messages are logged appropriately.

---

# 5. Error Handling and Logging Integration

    Error Aggregation:
        If any module fails, its error messages are recorded in its corresponding TestResult and included in the aggregated output.
    Logging Data:
        The Logging Module provides a chronological list of log entries that are appended to the JSON report.
    Resilience:
        The Results Generator is designed to continue aggregating data even if one or more modules report errors, ensuring that partial results can still be reviewed.
