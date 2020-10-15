package main

import (
	"flag"
	"os"
)

func main() {
	jsonFile := getCommandLineFlags()
	openFile(*jsonFile)
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
