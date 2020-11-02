package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"example.com/linkparser"
)

func main() {
	urlStr := getCommandLineFlags()
	links := bfs(*urlStr)

	for _, link := range links {
		fmt.Println(link)
	}
}

func getCommandLineFlags() *string {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	flag.Parse()
	return urlFlag
}

func bfs(urlStr string) []string {
	visited := make(map[string]struct{})
	queue := map[string]struct{}{
		urlStr: struct{}{},
	}

	for 0 < len(queue) {
		nextQueue := make(map[string]struct{})
		for key := range queue {
			if _, ok := visited[key]; !ok {
				visited[key] = struct{}{}
				for _, link := range getAllLinks(key) {
					nextQueue[link] = struct{}{}
				}
			}
		}
		queue = nextQueue
	}

	links := make([]string, 0, len(visited))
	for link := range visited {
		links = append(links, link)
	}
	return links
}

func getAllLinks(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	base := baseURL.String()
	return filter(getLinks(resp.Body, base), withPrefix(base))
}

func getLinks(r io.Reader, base string) []string {
	links, _ := linkparser.Parse(r)

	var formattedLinks []string
	for _, link := range links {
		switch {
		case strings.HasPrefix(link.Href, "/"):
			formattedLinks = append(formattedLinks, base+link.Href)
		case strings.HasPrefix(link.Href, "http"):
			formattedLinks = append(formattedLinks, link.Href)
		}
	}
	return formattedLinks
}

func filter(links []string, keepFunc func(string) bool) []string {
	var result []string

	for _, link := range links {
		if keepFunc(link) {
			result = append(result, link)
		}
	}

	return result
}

func withPrefix(prefix string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, prefix)
	}
}
