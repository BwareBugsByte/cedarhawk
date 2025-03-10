package ui

import (
	"cedarhawk/internal/config"
	"testing"
)

func TestRunResponsivenessTests_Pass(t *testing.T) {
	// Simulate HTML that includes responsive markers for desktop, tablet, and mobile.
	htmlData := `
	<html>
		<head>
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
		</head>
		<body>
			<div class="desktop-responsive">Desktop Layout</div>
			<div class="tablet-responsive">Tablet Layout</div>
			<div class="mobile-responsive">Mobile Layout</div>
		</body>
	</html>
	`

	// Define sample viewports.
	viewports := []config.Viewport{
		{Name: "desktop", Width: 1920, Height: 1080},
		{Name: "tablet", Width: 768, Height: 1024},
		{Name: "mobile", Width: 375, Height: 667},
	}

	resultsSlice := RunResponsivenessTests(htmlData, viewports)
	for _, res := range resultsSlice {
		if res.Status != "pass" {
			t.Errorf("Expected viewport %s to pass, got status '%s' with issues: %v", res.Viewport, res.Status, res.Issues)
		}
	}
}

func TestRunResponsivenessTests_Fail(t *testing.T) {
	// Simulate HTML without any responsive markers.
	htmlData := `
	<html>
		<body>
			<div>No responsive layout markers here.</div>
		</body>
	</html>
	`

	// Define sample viewports.
	viewports := []config.Viewport{
		{Name: "desktop", Width: 1920, Height: 1080},
		{Name: "tablet", Width: 768, Height: 1024},
		{Name: "mobile", Width: 375, Height: 667},
	}

	resultsSlice := RunResponsivenessTests(htmlData, viewports)
	for _, res := range resultsSlice {
		if res.Status != "fail" {
			t.Errorf("Expected viewport %s to fail, got status '%s'", res.Viewport, res.Status)
		}
		if len(res.Issues) == 0 {
			t.Errorf("Expected issues for viewport %s, but got none", res.Viewport)
		}
	}
}

func TestRunResponsivenessTests_Partial(t *testing.T) {
	// Simulate HTML that includes marker only for mobile.
	htmlData := `
	<html>
		<body>
			<div class="mobile-responsive">Mobile Layout</div>
		</body>
	</html>
	`

	// Define sample viewports.
	viewports := []config.Viewport{
		{Name: "desktop", Width: 1920, Height: 1080},
		{Name: "tablet", Width: 768, Height: 1024},
		{Name: "mobile", Width: 375, Height: 667},
	}

	resultsSlice := RunResponsivenessTests(htmlData, viewports)
	// Expect desktop and tablet to fail, mobile to pass.
	for _, res := range resultsSlice {
		if res.Viewport == "mobile" {
			if res.Status != "pass" {
				t.Errorf("Expected mobile to pass, got status '%s'", res.Status)
			}
		} else {
			if res.Status != "fail" {
				t.Errorf("Expected viewport %s to fail, got status '%s'", res.Viewport, res.Status)
			}
		}
	}
}

