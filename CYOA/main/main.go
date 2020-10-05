package main

import (
	"flag"
	"log"
	"net/http"
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

	h := story.NewHandler(s)
	log.Fatal(http.ListenAndServe(":3000", h))
}

func getCommandLineFlags() *string {
	jsonFlag := flag.String("file", "../gopher.json", "the JSON file with the CYOA story")
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
