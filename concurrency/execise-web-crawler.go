package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string, depth int, ch chan fetchResult)
}

type fetchResult struct {
	depth int
	url   string
	body  string
	urls  []string
	err   error
}

type SafeCache struct {
	m        sync.Mutex
	urlCache map[string]fetchResult
}

func (sc *SafeCache) isCached(url string) bool {
	_, ok := sc.urlCache[url]
	return ok
}
func (sc *SafeCache) cache(result fetchResult) {
	sc.m.Lock()
	sc.urlCache[result.url] = result
	sc.m.Unlock()
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.

func Crawl(url string, depth int, fetcher Fetcher) {
	sc := SafeCache{urlCache: make(map[string]fetchResult)}
	ch := make(chan fetchResult)
	crawlRecursive(url, depth, fetcher, ch)
	for {
		if result, ok := <-ch; ok {
			if sc.isCached(result.url){
				fmt.Printf("cached[%d]: %s %q\n", result.depth, result.url, result.body)
			} else if result.err != nil {
				fmt.Println(result.err)
				sc.cache(result)
			} else {
				fmt.Printf("found[%d]: %s %q\n", result.depth, result.url, result.body)
				sc.cache(result)
				for _, u := range result.urls {
					crawlRecursive(u, result.depth-1, fetcher, ch)
				}
			}
		} else {
			close(ch)
			return
		}
	}
}

func crawlRecursive(url string, depth int, fetcher Fetcher, ch chan fetchResult) {
	if depth <= 0 {
		return
	}

	go fetcher.Fetch(url, depth, ch)

}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string, depth int, ch chan fetchResult) {
	if res, ok := f[url]; ok {
		ch <- fetchResult{
			url:   url,
			body:  res.body,
			urls:  res.urls,
			err:   nil,
			depth: depth,
		}
		return
	}

	ch <- fetchResult{
		url:   url,
		body:  "",
		urls:  nil,
		err:   fmt.Errorf("not found: %s", url),
		depth: depth,
	}
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
