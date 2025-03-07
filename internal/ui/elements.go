package ui

import (
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"cedarhawk/internal/results"
)

// RunElementsValidation validates the presence and functionality of essential UI elements
// in the provided HTML data. It returns a slice of UIElementResult detailing the outcome
// for each element type.
func RunElementsValidation(htmlData string) []results.UIElementResult {
	var resultsSlice []results.UIElementResult

	// -------- Validate Buttons --------
	// Search for <button> elements and <input type="button"> elements.
	buttonRegex := regexp.MustCompile(`(?i)<button[^>]*>`)
	inputButtonRegex := regexp.MustCompile(`(?i)<input[^>]*type=["']button["'][^>]*>`)
	buttonMatches := buttonRegex.FindAllString(htmlData, -1)
	inputButtonMatches := inputButtonRegex.FindAllString(htmlData, -1)
	totalButtons := len(buttonMatches) + len(inputButtonMatches)

	buttonStatus := "pass"
	buttonMessage := "Buttons are present."
	if totalButtons == 0 {
		buttonStatus = "fail"
		buttonMessage = "No buttons found."
		log.Println("UI Elements Validation: No buttons found.")
	} else {
		// Optionally, check that <button> tags include an "onclick" attribute.
		missingOnclick := false
		for _, btn := range buttonMatches {
			if !strings.Contains(strings.ToLower(btn), "onclick=") {
				missingOnclick = true
				break
			}
		}
		if missingOnclick {
			buttonMessage = "Buttons found but some may be missing 'onclick' attributes."
			log.Println("UI Elements Validation: Some buttons are missing 'onclick' attribute.")
		}
	}
	resultsSlice = append(resultsSlice, results.UIElementResult{
		ElementType: "button",
		Status:      buttonStatus,
		Message:     buttonMessage,
	})

	// -------- Validate Links --------
	linkRegex := regexp.MustCompile(`(?i)<a\s+[^>]*href=["']([^"']+)["']`)
	linkMatches := linkRegex.FindAllStringSubmatch(htmlData, -1)
	linkStatus := "pass"
	linkMessage := "Links are present."
	if len(linkMatches) == 0 {
		linkStatus = "fail"
		linkMessage = "No links found."
		log.Println("UI Elements Validation: No links found.")
	}
	resultsSlice = append(resultsSlice, results.UIElementResult{
		ElementType: "link",
		Status:      linkStatus,
		Message:     linkMessage,
	})

	// -------- Validate Images --------
	imgRegex := regexp.MustCompile(`(?i)<img\s+[^>]*src=["']([^"']+)["']`)
	imgMatches := imgRegex.FindAllStringSubmatch(htmlData, -1)
	imageStatus := "pass"
	imageMessage := "All images loaded successfully."
	if len(imgMatches) == 0 {
		imageStatus = "fail"
		imageMessage = "No images found."
		log.Println("UI Elements Validation: No images found.")
	} else {
		// For each image, perform an HTTP GET request to check if the image is accessible.
		client := &http.Client{
			Timeout: 5 * time.Second,
		}
		for _, match := range imgMatches {
			if len(match) < 2 {
				continue
			}
			src := match[1]
			resp, err := client.Get(src)
			if err != nil || resp.StatusCode != http.StatusOK {
				imageStatus = "fail"
				imageMessage = "One or more images failed to load. Problematic image src: " + src
				log.Printf("UI Elements Validation: Image check failed for %s: %v", src, err)
				if resp != nil {
					resp.Body.Close()
				}
				break
			}
			resp.Body.Close()
		}
	}
	resultsSlice = append(resultsSlice, results.UIElementResult{
		ElementType: "image",
		Status:      imageStatus,
		Message:     imageMessage,
	})

	// -------- Validate Text Fields --------
	textInputRegex := regexp.MustCompile(`(?i)<input\s+[^>]*type=["']text["'][^>]*>`)
	textareaRegex := regexp.MustCompile(`(?i)<textarea[^>]*>`)
	textInputs := textInputRegex.FindAllString(htmlData, -1)
	textareas := textareaRegex.FindAllString(htmlData, -1)
	totalTextFields := len(textInputs) + len(textareas)
	textStatus := "pass"
	textMessage := "Text fields are present."
	if totalTextFields == 0 {
		textStatus = "fail"
		textMessage = "No text fields found."
		log.Println("UI Elements Validation: No text fields found.")
	}
	resultsSlice = append(resultsSlice, results.UIElementResult{
		ElementType: "textField",
		Status:      textStatus,
		Message:     textMessage,
	})

	return resultsSlice
}
