package ui

import (
	"strings"
	"testing"
)

func TestRunLayoutVerification_Pass(t *testing.T) {
	// HTML that includes all required elements for the default layout.
	htmlData := `
	<html>
		<header>This is the header</header>
		<nav>This is navigation</nav>
		<main>This is main content</main>
		<footer>This is the footer</footer>
	</html>
	`
	result := RunLayoutVerification(htmlData, "default")
	if result.Status != "pass" {
		t.Errorf("Expected status 'pass', got '%s'. Message: %s", result.Status, result.Message)
	}
}

func TestRunLayoutVerification_Fail(t *testing.T) {
	// HTML missing the <nav> and <footer> elements.
	htmlData := `
	<html>
		<header>This is the header</header>
		<main>This is main content</main>
	</html>
	`
	result := RunLayoutVerification(htmlData, "default")
	if result.Status != "fail" {
		t.Errorf("Expected status 'fail', got '%s'", result.Status)
	}
	// Check that the error message mentions the missing <nav> and <footer> elements.
	if !strings.Contains(result.Message, "<nav") || !strings.Contains(result.Message, "<footer") {
		t.Errorf("Expected missing elements message, got: %s", result.Message)
	}
}

func TestRunLayoutVerification_CustomLayout(t *testing.T) {
	// For a custom layout, define expected elements differently.
	htmlData := `
	<html>
		<div id="content">Content here</div>
		<p>Some text</p>
	</html>
	`
	result := RunLayoutVerification(htmlData, "custom")
	// With our placeholder custom layout, we expect to find <div and <p.
	if result.Status != "pass" {
		t.Errorf("Expected status 'pass' for custom layout, got '%s'. Message: %s", result.Status, result.Message)
	}
}

