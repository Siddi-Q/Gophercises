package main

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := MapHandler(pathsToUrls, mux)

	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	yamlHandler, err := YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

// MapHandler will return an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths (keys in the map) to their corresponding URL
// (values that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return a http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding URL
// If the path is not provided in the YAML, then the fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
// - path: /some-path
//   url: https://www.some-url.com/demo
//
// The only errors that can be returned are all related to having invalid YAML data.
func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathURLs, err := parseYaml(yamlBytes)
	if err != nil {
		return nil, err
	}

	pathsToUrls := buildMap(pathURLs)
	return MapHandler(pathsToUrls, fallback), nil
}

func parseYaml(yamlBytes []byte) ([]pathURL, error) {
	var pathURLs []pathURL
	err := yaml.Unmarshal(yamlBytes, &pathURLs)

	if err != nil {
		return nil, err
	}
	return pathURLs, nil
}

func buildMap(pathURLs []pathURL) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, pathURL := range pathURLs {
		pathsToUrls[pathURL.Path] = pathURL.URL
	}
	return pathsToUrls
}

type pathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
