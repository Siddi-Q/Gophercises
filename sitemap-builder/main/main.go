package main

import (
	"flag"
	"fmt"
)

func main() {
	url := getCommandLineFlags()
	fmt.Println(*url)
}

func getCommandLineFlags() *string {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	flag.Parse()
	return urlFlag
}
