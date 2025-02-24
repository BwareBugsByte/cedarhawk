package crawler

import (
	"log"
)

// Crawl starts at the given URL and recursively fetches HTML pages.
func Crawl(url string, depth int) ([]string, error) {
	// TODO: Implement crawling logic using net/http and avoid loops via a visited map.
	log.Println("Placeholder: Starting crawl at", url)
	return []string{url}, nil
}
