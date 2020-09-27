package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"handler"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := handler.MapHandler(pathsToUrls, mux)

	yamlFile := getCommandLineFlags()
	yaml := readYamlFile(*yamlFile)
	yamlHandler, err := handler.YAMLHandler(yaml, mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func getCommandLineFlags() *string {
	yamlFlag := flag.String("yaml", "../pathToUrls.yaml", "a yaml file in the format of\n '- path: /insertPathHere \n    url: https://insertUrlHere'")
	flag.Parse()
	return yamlFlag
}

func readYamlFile(yamlFile string) []byte {
	f, err := os.Open(yamlFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return b
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
