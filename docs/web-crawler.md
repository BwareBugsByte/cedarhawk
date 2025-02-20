Design: Crawler Module
================
# 1. Overview

The Crawler Module is responsible for recursively traversing a website, starting from a given base URL. It gathers HTML pages for further analysis by other modules while avoiding loops and duplicate URL processing. This module is implemented using only Go’s built-in libraries.

---

# 2. Core Components
Entry Point

    Function: Initiates the crawl from a configured base URL.
    Invocation: Called from the main entry point of the application.

Visited Map

    Purpose: Maintains a record of all URLs that have been visited.
    Implementation: A map[string]bool that is checked before processing a new URL.
    Benefit: Prevents infinite loops and duplicate processing.

HTTP Fetcher

    Function: Uses net/http to fetch HTML content.
    Implementation:
        Use http.Get(url) to retrieve page content.
        Implement timeouts using an http.Client with a custom timeout setting.
    Error Handling: Log and skip URLs that fail to fetch.

HTML Parser

    Function: Extracts links from the fetched HTML.
    Implementation:
        Use simple string operations or regular expressions to search for <a href="..."> tags.
        Use the net/url package to resolve relative URLs.
    Limitation: The parser is designed for basic HTML structures; advanced HTML parsing is out of scope to maintain built-in dependency restrictions.

Recursive or Queue-Based Crawler

    Mechanism:
        Recursion: The crawler calls itself with each newly discovered URL, decrementing a depth counter.
        Queue: Alternatively, use a loop with a FIFO queue for breadth-first traversal.
    Configuration: A maximum recursion depth (e.g., maxDepth from configuration) is used to limit crawl depth.

---

# 3. Avoiding Loops and Duplicate URLs

    Visited Map:
        Every URL is checked against the visited map before processing.
        If a URL is already marked as visited, it is skipped.
    Domain Constraints:
        Optionally, restrict crawling to URLs within the same domain as the base URL.

---

# 4. Fetching and Parsing HTML
Fetching HTML

    Method:
        Use the built-in net/http package.
        Example:

        resp, err := http.Get(url)
        if err != nil {
            // Log error and return
        }
        defer resp.Body.Close()

    Timeouts:
        Implement using http.Client{Timeout: time.Second * 10} to avoid hanging requests.

Parsing HTML

    Approach:
        Use regular expressions or string functions to extract <a> tags.
        Example (simplified):

        // Pseudocode: extract href attributes from <a> tags
        hrefPattern := regexp.MustCompile(`(?i)<a\s+(?:[^>]*?\s+)?href="([^"]*)"`)
        matches := hrefPattern.FindAllStringSubmatch(htmlContent, -1)

    URL Resolution:
        Use net/url.Parse and ResolveReference to handle relative URLs.

5. Error Handling and Performance Considerations
Error Handling

    HTTP Errors:
        Log errors if an HTTP request fails and continue with the next URL.
    Parsing Errors:
        If parsing fails, log the error along with the problematic URL.
    Critical Errors:
        For configuration or network-level critical errors, propagate the error to terminate the crawl gracefully.
    Error Reporting:
        Errors are recorded and later aggregated in the JSON results.

Performance Considerations

    Concurrency:
        Optionally, implement concurrent crawling using goroutines.
        Use synchronization (e.g., a mutex) when updating the visited map.
    Rate Limiting:
        Implement delays between requests to avoid overwhelming target servers.
    Resource Cleanup:
        Ensure HTTP response bodies are closed promptly to free resources.

---

# 6. Flowchart and Sequence Diagram
Flowchart:

           +-----------------------+
           | Start with Base URL   |
           +-----------+-----------+
                       |
                       v
           +-----------------------+
           | Check if URL visited? |
           +-----------+-----------+
                       | Yes -> Skip URL
                       v No
           +-----------------------+
           | Mark URL as visited   |
           +-----------+-----------+
                       |
                       v
           +-----------------------+
           | Fetch HTML Content    |
           | (http.Get)            |
           +-----------+-----------+
                       |
                       v
           +-----------------------+
           | Parse HTML to extract |
           | links (href attributes)|
           +-----------+-----------+
                       |
                       v
           +-----------------------+
           | For each extracted    |
           | link, validate & add  |
           | to crawl if not visited|
           +-----------+-----------+
                       |
                       v
           +-----------------------+
           | Recursively process   |
           | next URL (or enqueue) |
           +-----------------------+

Sequence Diagram:

User
  |
  v
Main Entry Point
  |-- calls --> Crawl(baseURL, maxDepth)
                       |
                       v
           +-----------------------+
           | Check Visited Map     |
           +-----------------------+
                       |
                       v
           +-----------------------+
           | HTTP GET (net/http)   |
           +-----------------------+
                       |
                       v
           +-----------------------+
           | Return HTML Content   |
           +-----------------------+
                       |
                       v
           +-----------------------+
           | Parse HTML for links  |
           +-----------------------+
                       |
                       v
           +-----------------------+
           | For each link:        |
           |   If not visited,     |
           |   call Crawl(link, d-1)|
           +-----------------------+

---

# 7. Handling Edge Cases

    Invalid URLs:
        Skip and log any URLs that fail to parse.
    Circular Links:
        Handled by the visited map.
    HTTP Request Failures:
        Log the error and move on to the next URL.
    Large or Deep Site Structures:
        Limit recursion with a maximum depth.
        Consider a timeout or maximum number of pages to avoid resource exhaustion.

---

# 8. Conclusion

This design for the Crawler Module outlines a robust and modular approach to recursively traverse a website starting from a base URL. It uses a visited map to prevent loops, leverages built-in Go libraries (net/http and standard string operations) for fetching and parsing HTML, and incorporates comprehensive error handling and performance strategies. The provided flowchart and sequence diagram illustrate the module's operational flow and help ensure clarity in implementation.
