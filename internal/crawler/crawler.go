package crawler

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"sync"
)

// PageData holds the URL and its fetched HTML content.
type PageData struct {
	URL  string
	HTML string
}

// Crawl recursively visits pages starting at startURL up to maxDepth.
// It returns a slice of PageData for each unique visited URL.
func Crawl(startURL string, maxDepth int) ([]PageData, error) {
	visited := make(map[string]bool)
	var result []PageData
	var mu sync.Mutex

	// Define a recursive function.
	var crawl func(string, int)
	crawl = func(currentURL string, depth int) {
		if depth < 0 {
			return
		}

		// Avoid duplicate URLs.
		mu.Lock()
		if visited[currentURL] {
			mu.Unlock()
			return
		}
		visited[currentURL] = true
		mu.Unlock()

		// Fetch the HTML page.
		resp, err := http.Get(currentURL)
		if err != nil {
			log.Printf("Error fetching URL %s: %v", currentURL, err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("Non-OK HTTP status for %s: %s", currentURL, resp.Status)
			return
		}

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading body for %s: %v", currentURL, err)
			return
		}
		bodyStr := string(bodyBytes)

		// Save the page data.
		mu.Lock()
		result = append(result, PageData{URL: currentURL, HTML: bodyStr})
		mu.Unlock()

		// Extract links using a simple regex.
		re := regexp.MustCompile(`(?i)<a\s+(?:[^>]*?\s+)?href="([^"]*)"`)
		matches := re.FindAllStringSubmatch(bodyStr, -1)
		for _, match := range matches {
			if len(match) < 2 {
				continue
			}
			link := match[1]

			// Parse and resolve the URL.
			parsedLink, err := url.Parse(link)
			if err != nil {
				log.Printf("Error parsing URL %s found on %s: %v", link, currentURL, err)
				continue
			}
			base, err := url.Parse(currentURL)
			if err != nil {
				log.Printf("Error parsing base URL %s: %v", currentURL, err)
				continue
			}
			absoluteURL := base.ResolveReference(parsedLink).String()

			// Recursively crawl the linked page.
			crawl(absoluteURL, depth-1)
		}
	}

	crawl(startURL, maxDepth)
	return result, nil
}
