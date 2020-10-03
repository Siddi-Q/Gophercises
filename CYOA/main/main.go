package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	jsonFile := getCommandLineFlags()
	s := parseJSON(*jsonFile)
	fmt.Printf("%+v", s)
}

func getCommandLineFlags() *string {
	jsonFlag := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	return jsonFlag
}

func parseJSON(jsonFile string) story {
	f, err := os.Open(jsonFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	d := json.NewDecoder(f)
	var s story
	if err := d.Decode(&s); err != nil {
		panic(err)
	}
	return s
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
