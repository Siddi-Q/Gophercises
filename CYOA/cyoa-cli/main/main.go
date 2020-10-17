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

	play(s)
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

func play(s story.Story) {
	chapterName := "intro"

	for len(s[chapterName].Options) != 0 {
		fmt.Printf("%v\n", s[chapterName].Title)
		for _, paragraph := range s[chapterName].Paragraphs {
			fmt.Printf("%v\n\n", paragraph)
		}

		for _, option := range s[chapterName].Options {
			fmt.Printf("%v: %v\n", option.Chapter, option.Text)
		}

		fmt.Println()

		for idx, option := range s[chapterName].Options {
			fmt.Printf("Press %v to venture to %v\n", idx+1, option.Chapter)
		}

		var choice int
		fmt.Scanf("%d\n", &choice)
		chapterName = s[chapterName].Options[choice-1].Chapter
	}
}
