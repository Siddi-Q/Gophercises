package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"example.com/story"
)

func main() {
	jsonFile, port := getCommandLineFlags()
	f := openFile(*jsonFile)
	s, err := story.ParseJSONStory(f)
	if err != nil {
		panic(err)
	}
	f.Close()

	h := story.NewHandler(s)
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}

func getCommandLineFlags() (*string, *int) {
	jsonFlag := flag.String("file", "../../gopher.json", "the JSON file with the CYOA story")
	portFlag := flag.Int("port", 3000, "the port to start the CYOA web application on")
	flag.Parse()
	return jsonFlag, portFlag
}

func openFile(fileName string) *os.File {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	return f
}
