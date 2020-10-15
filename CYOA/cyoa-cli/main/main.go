package main

import (
	"flag"
	"fmt"
	"os"

	"example.com/story"
)

func main() {
	jsonFile := getCommandLineFlags()
	f := openFile(*jsonFile)
	s, err := story.ParseJSONStory(f)
	if err != nil {
		panic(err)
	}
	f.Close()
	fmt.Printf("%+v", s)
}

func getCommandLineFlags() *string {
	jsonFlag := flag.String("file", "../../gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	return jsonFlag
}

func openFile(fileName string) *os.File {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	return f
}
