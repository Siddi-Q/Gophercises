package main

import "flag"

func main() {
	jsonFile := getCommandLineFlags()
}

func getCommandLineFlags() *string {
	jsonFlag := flag.String("file", "../gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	return jsonFlag
}
