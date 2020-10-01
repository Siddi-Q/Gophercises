package main

import "flag"

func main() {
}

func getCommandLineFlags() *string {
	jsonFlag := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	return jsonFlag
}

type story map[string]chapter

type chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []option `json:"options"`
}

type option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
