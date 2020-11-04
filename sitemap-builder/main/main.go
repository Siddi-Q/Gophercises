package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"example.com/linkparser"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {
	urlStr, _ := getCommandLineFlags()
	pages := bfs(*urlStr)
	toXML(pages)
}

func getCommandLineFlags() (*string, *int) {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	maxDepthFlag := flag.Int("depth", 10, "the maximum number of links deep to traverse")
	flag.Parse()
	return urlFlag, maxDepthFlag
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
		return []string{}
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

func toXML(pages []string) {
	sitemap := urlset{Xmlns: xmlns}
	for _, page := range pages {
		sitemap.Urls = append(sitemap.Urls, loc{page})
	}

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")

	fmt.Print(xml.Header)
	if err := enc.Encode(sitemap); err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println()
}
