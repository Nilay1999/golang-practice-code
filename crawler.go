package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchUrl(url string) string {
	time.Sleep(time.Second)
	return "Content fetched from " + url
}

func crawl(urls []string) map[string]string {
	urlChannel := make(chan [2]string) // Channel to send URL and fetched content pairs
	result := make(map[string]string)
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			content := fetchUrl(url)
			urlChannel <- [2]string{url, content}
		}(url) // Pass url as argument to the anonymous function
	}

	// Goroutine to close the channel after all goroutines complete
	go func() {
		wg.Wait()
		close(urlChannel)
	}()

	// Collect results from the channel
	for pair := range urlChannel {
		result[pair[0]] = pair[1]
	}

	return result
}

func main() {
	url := [5]string{"http://example.com", "http://example.org", "http://example.net", "http://example.edu", "http://example.io"}
	result := crawl(url[:])
	for key := range result {
		fmt.Printf("%s\n", key)
	}
}
