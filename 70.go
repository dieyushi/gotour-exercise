/*
Exercise: Web Crawler

In this exercise you'll use Go's concurrency features to parallelize a web crawler.

Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice.
*/

package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	type target struct {
		urls  []string
		depth int
	}
	more := make(chan target)

	fetchPage := func(url string, depth int) {
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("found: %s %q\n", url, body)
		}
		more <- target{urls, depth + 1}
	}

	go fetchPage(url, 0)

	visited := map[string]bool{url: true}
	for counter := 1; counter > 0; {
		next := <-more
		counter--
		if next.depth > depth {
			continue
		}
		for _, url := range next.urls {
			if _, seen := visited[url]; seen {
				continue
			}
			visited[url] = true
			counter++
			go fetchPage(url, next.depth)
		}
	}
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := (*f)[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
