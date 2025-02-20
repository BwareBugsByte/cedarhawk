High-Level Data Flow Diagram
====
# Below is an ASCII diagram illustrating the flow of data and interactions between components:

                +----------------------+
                |  Main Entry Point    |
                +----------+-----------+
                           |
         +-----------------+-----------------+
         |                                   |
+--------v---------+                +--------v---------+
| Configuration    |                | Logging & Error  |
| Loader           |                | Handling Module  |
+--------+---------+                +--------+---------+
         |                                   |
         v                                   |
+--------+---------+                         |
|  Crawler Module  |<------------------------+
| (Fetch & Parse   |
| HTML Pages)      |
+--------+---------+
         |
         v
+--------+---------+   +----------------+   +----------------+
| Layout           |   | UI Elements    |   | Responsiveness |
| Verification     |   | Validation     |   | Testing Module |
| Module           |   | Module         |   |                |
+--------+---------+   +----------------+   +----------------+
         |                         |                |
         +------------+------------+----------------+
                      |
                      v
            +---------+----------+
            |  Results Generator |
            | (JSON Output)      |
            +---------+----------+
                      |
                      v
          [Optional AI Plugin Interface]

# Key Design Decisions, Assumptions, and Constraints

    Built-in Libraries Only:
        The design leverages only Go’s built-in packages (e.g., net/http, encoding/json) to minimize external dependencies.
    Modularity:
        Each component is designed as an independent module, making it easy to test, maintain, and potentially replace.
    Configuration-Driven:
        The behavior of CedarHawk is governed by the user-provided config.toml, allowing flexible customization without code changes.
    Error Handling and Logging:
        Centralized logging and error handling are integrated across modules, ensuring that failures are captured and reported consistently.
    Future-Proofing with AI Plugin:
        The architecture reserves a clear interface for future AI enhancements without impacting the core functionalities.
