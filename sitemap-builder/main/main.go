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

	links := getAllLinks(*urlStr)
	for _, link := range links {
		fmt.Println(link)
	}
}

func getCommandLineFlags() *string {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	flag.Parse()
	return urlFlag
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
