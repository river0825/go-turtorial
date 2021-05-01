package main

import "testing"

func TestCrawl(t *testing.T) {
	type args struct {
		url     string
		depth   int
		fetcher Fetcher
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Crawl main", args: args{url: "https://golang.org/", depth: 4, fetcher: fetcher}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Crawl("https://golang.org/", 4, fetcher)
		})
	}
}
