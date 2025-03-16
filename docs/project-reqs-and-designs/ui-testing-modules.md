Design: UI Testing Modules
=====================
The UI Testing Modules are a core part of CedarHawk and are responsible for ensuring that the website's user interface behaves as expected. These modules work together to verify layout consistency, validate the presence and functionality of essential UI elements, and confirm that the site is responsive across various viewport sizes.
a. Layout Verification Module
Role

    DOM Extraction:
        Parse the HTML provided by the Crawler Module to build a representation of the page’s DOM. This includes identifying the container hierarchy and common elements such as headers, footers, sidebars, and content areas.
    Layout Comparison:
        Compare the extracted DOM structure against expected layouts defined in the configuration file.
        Check for consistency across different pages and screen sizes.
        Identify any discrepancies in layout, such as missing or misaligned components.

Interactions

    Input:
        Receives HTML/DOM data from the Crawler Module.
        Uses configuration parameters (from the Configuration Loader) that specify expected layout patterns and key elements.
    Processing:
        Extracts and represents the layout in a structured format (e.g., tree or graph).
        Compares the layout structure across multiple pages and viewports.
    Output:
        Generates a JSON report highlighting any layout inconsistencies, discrepancies, or missing elements.

b. UI Elements Validation Module
Role

    Element Presence Check:
        Verify that essential UI components (buttons, links, images, text fields) are present on each page.
    Functionality Testing:
        Confirm that UI elements function as intended. For example, validate that clickable elements have the correct attributes (e.g., onclick handlers).
    Image Verification:
        Check that images load successfully by validating that image URLs return a proper response and are not broken.

Interactions

    Input:
        Processes HTML data provided by the Crawler Module.
        Utilizes validation rules defined in the configuration file for determining which elements are required and any specific attributes they must possess.
    Processing:
        Scans the DOM for required elements and inspects their attributes.
        Performs HTTP checks on image URLs to ensure they return valid responses.
    Output:
        Produces a detailed JSON report summarizing the status of each UI element, including any missing elements, malfunctioning components, or broken links.

c. Responsiveness Testing Module
Role

    Viewport Simulation:
        Simulate various viewport sizes (e.g., desktop, tablet, mobile) based on settings provided in the configuration file.
    Layout Adaptation:
        Verify that the UI adjusts appropriately to different viewport sizes. This includes checking that elements scale correctly, remain accessible, and that the overall layout does not break.
    Responsive Behavior Analysis:
        Evaluate how dynamic components behave under different viewport constraints.

Interactions

    Input:
        Receives viewport configuration settings from the Configuration Loader.
        Uses HTML/DOM data from the Crawler Module.
    Processing:
        For each defined viewport size, simulate the environment and analyze the DOM to detect any layout shifts or issues.
        Compare actual UI element positioning and scaling against expected outcomes.
    Output:
        Generates a JSON report that details the responsiveness test results for each viewport size, including any detected issues or layout breaks.

Summary Diagram of UI Testing Modules

                +----------------------+
                |  Configuration       |
                |      Loader          |
                +----------+-----------+
                           |
                           v
                +----------------------+
                |  Crawler Module      |
                | (HTML/DOM Extraction)|
                +----------+-----------+
                           |
                           v
             +-------------+-------------+
             |   UI Testing Modules      |
             |  +---------+  +---------+  |
             |  | Layout  |  | UI      |  |
             |  | Verify  |  | Elements|  |
             |  | Module  |  | Validation|
             |  +---------+  +---------+  |
             |       +-----------------+ |
             |       | Responsiveness  | |
             |       | Testing Module  | |
             |       +-----------------+ |
             +-------------+-------------+
                           |
                           v
                +----------------------+
                |  Results Generator   |
                +----------------------+

---

Design: UI Testing Module (UI Elements Validation)
=========================================
# 1. Overview

The UI Elements Validation module is responsible for ensuring that essential UI components are present and functioning correctly on a webpage. This includes verifying the existence and behavior of buttons, links, images, and text fields. Additionally, it will detect broken image links and, optionally, validate interactive attributes (e.g., onclick handlers) where specified.

---

# 2. Responsibilities

    Presence Check:
        Ensure that required UI elements (buttons, links, images, text fields) are present on the page.
    Functionality Check:
        Verify that interactive elements (e.g., buttons and links) are correctly configured. For example, check that buttons have an onclick attribute when required.
    Image Verification:
        Validate that images load correctly by performing a simple HTTP check on the image source URL using the built-in net/http package.
    Reporting:
        Generate detailed output for each check, indicating success or failure along with relevant error messages.

---

# 3. Architecture and Data Flow
Component Interactions

    Input Data:
        Receives pre-parsed HTML/DOM data from the Crawler Module.
        Retrieves validation rules and criteria from the configuration (loaded by the Configuration Loader).

    Processing:
        Iterates over the DOM to locate UI elements.
        For each element type:
            Buttons: Check for existence and inspect interactive attributes like onclick if required.
            Links: Verify that a valid href attribute is present.
            Images: Confirm that each image element has a valid src attribute. Perform an HTTP GET request to detect broken links.
            Text Fields: Ensure text fields are present and optionally check for expected attributes (e.g., placeholder text).

    Output Data:
        Compiles results into a structured JSON format indicating:
            Element type
            Check status (pass/fail)
            Error details (e.g., "Missing element," "Broken link," "Missing onclick attribute")
            Additional metadata (e.g., element identifier or location)

Data Flow Diagram

                +----------------------+
                |  Crawler Module      |
                | (HTML/DOM Extraction)|
                +----------+-----------+
                           |
                           v
              +------------+------------+
              | UI Elements Validation  |
              |        Module           |
              +------------+------------+
                           |
                           v
              +------------+------------+
              |   JSON Results Output   |
              +-------------------------+

---

# 4. Implementation Strategy
Step-by-Step Process

    Element Extraction:
        Use DOM traversal (e.g., via Go's golang.org/x/net/html or a built-in parser approach) to iterate through the HTML.
    Validation Checks:
        Buttons:
            Confirm at least one button is present.
            Optionally, if the configuration specifies, ensure the button has an onclick attribute.
        Links:
            Verify that links have a non-empty href attribute.
        Images:
            For each <img> element, check for the src attribute.
            Use net/http to perform a GET request on the image URL to ensure it returns a valid response (e.g., status code 200).
        Text Fields:
            Ensure that input fields of type text (or similar) are present.
    Error Handling:
        For each check, if the expected condition is not met, record a detailed error message.
        Handle HTTP errors for images gracefully, marking an image as broken if the GET request fails or returns a non-success status code.
        Aggregate all error messages with clear identifiers (e.g., element type, index, or CSS selector) to help pinpoint issues.

---

# 5. Expected Behavior

    Successful Validation:
        The module reports a “pass” status for each element type if all validations succeed.
        A summary message indicates that all required UI elements were found and function as expected.
    Failure Cases:
        Missing Elements: Report which required UI element (button, link, image, text field) is absent.
        Broken Links/Images: Report the URL and error received (e.g., HTTP status code) for any broken image link.
        Missing Interactive Attributes: If enabled in configuration, report elements (such as buttons) that are missing expected interactive attributes like onclick.
        Each error is documented with its location (e.g., page URL or element identifier) to facilitate troubleshooting.

---

# 6. Error Reporting Strategy

    Structured JSON Output:
        The module will produce a JSON report with an array of check results. Each entry includes:
            elementType: "button", "link", "image", "textField"
            status: "pass" or "fail"
            message: Detailed error message if the check fails.
            elementIdentifier: (Optional) Specific identifier or index for the element.

    Example JSON Snippet:

{
  "page": "https://example.com",
  "uiElements": [
    {
      "elementType": "button",
      "status": "pass",
      "message": "Button found with valid onclick attribute.",
      "elementIdentifier": "btn-submit"
    },
    {
      "elementType": "image",
      "status": "fail",
      "message": "Image source returned HTTP 404: https://example.com/img/logo.png",
      "elementIdentifier": "img-logo"
    },
    {
      "elementType": "link",
      "status": "pass",
      "message": "Link contains valid href attribute.",
      "elementIdentifier": "a-home"
    }
  ]
}

    Error Logging:
        In addition to the JSON output, errors will be logged using the Logging Module, allowing for real-time diagnostics during test execution.

---

# 7. Sample Checks and Strategies

    Button Check:
        Check: Ensure at least one <button> or <input type="button"> is present.
        Optional: If required by config, check that the button contains an onclick attribute.
        Error: "No button found" or "Button missing onclick attribute."

    Link Check:
        Check: Validate that <a> tags have a non-empty href attribute.
        Error: "Link with empty href found."

    Image Check:
        Check: Verify <img> tags have a src attribute.
        HTTP Check: Use a GET request on the src URL and check for a 200 response.
        Error: "Broken image detected with src: [URL] (HTTP 404)."

    Text Field Check:
        Check: Validate presence of <input> elements with type text or <textarea>.
        Error: "Required text field missing."
