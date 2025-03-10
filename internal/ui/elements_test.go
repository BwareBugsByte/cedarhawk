package ui

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRunElementsValidation_AllPresent(t *testing.T) {
	// Create a test HTTP server to simulate image endpoints.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	// HTML including:
	// - A <button> with an onclick attribute,
	// - A link,
	// - An image with a valid src,
	// - A text input and a textarea.
	htmlData := `
	<html>
		<body>
			<button onclick="alert('Clicked!')">Click me</button>
			<a href="https://example.com">Example Link</a>
			<img src="` + ts.URL + `/image.png" alt="Test Image">
			<input type="text" name="username">
			<textarea name="comments"></textarea>
		</body>
	</html>
	`
	resultsSlice := RunElementsValidation(htmlData)
	// All element validations should pass.
	for _, res := range resultsSlice {
		if res.Status != "pass" {
			t.Errorf("Expected element type %s to pass, got status '%s' with message: %s", res.ElementType, res.Status, res.Message)
		}
	}
}

func TestRunElementsValidation_MissingElements(t *testing.T) {
	// HTML missing buttons, images, and text fields; only includes a link.
	htmlData := `
	<html>
		<body>
			<a href="https://example.com">Example Link</a>
		</body>
	</html>
	`
	resultsSlice := RunElementsValidation(htmlData)
	var buttonFound, imageFound, textFieldFound bool
	for _, res := range resultsSlice {
		switch res.ElementType {
		case "button":
			buttonFound = true
			if res.Status != "fail" {
				t.Errorf("Expected button status 'fail', got '%s'", res.Status)
			}
		case "image":
			imageFound = true
			if res.Status != "fail" {
				t.Errorf("Expected image status 'fail', got '%s'", res.Status)
			}
		case "textField":
			textFieldFound = true
			if res.Status != "fail" {
				t.Errorf("Expected textField status 'fail', got '%s'", res.Status)
			}
		}
	}
	if !buttonFound {
		t.Error("Button result not found in UI Elements Validation")
	}
	if !imageFound {
		t.Error("Image result not found in UI Elements Validation")
	}
	if !textFieldFound {
		t.Error("TextField result not found in UI Elements Validation")
	}
}

func TestRunElementsValidation_BrokenImage(t *testing.T) {
	// HTML with an image using an unreachable URL.
	htmlData := `
	<html>
		<body>
			<img src="http://nonexistent.example.com/image.png" alt="Broken Image">
		</body>
	</html>
	`
	resultsSlice := RunElementsValidation(htmlData)
	for _, res := range resultsSlice {
		if res.ElementType == "image" {
			if res.Status != "fail" {
				t.Errorf("Expected image status 'fail' for broken image, got '%s'", res.Status)
			}
			if !strings.Contains(strings.ToLower(res.Message), "failed") {
				t.Errorf("Expected error message for broken image, got '%s'", res.Message)
			}
		}
	}
}

