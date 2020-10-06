package story

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

var defaultHTMLTemplate = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>Choose Your Own Adventure</title>
	</head>
	<body>
		<h1>{{.Title}}</h1>
		{{range .Paragraphs}}
			<p>{{.}}</p>
		{{end}}
		<ul>
		{{range .Options}}
			<li><a href="/{{.Chapter}}">{{.Text}}</a></li>
		{{end}}
		</ul>
	</body>
</html>
`

var t = template.Must(template.New("").Parse(defaultHTMLTemplate))

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)

	if path == "" || path == "/" {
		path = "/intro"
	}

	path = path[1:] // removes the '/' from path

	if chapter, ok := h.s[path]; ok {
		err := t.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found.", http.StatusNotFound)
}

// NewHandler will
func NewHandler(s Story) http.Handler {
	return handler{s}
}

// ParseJSONStory will
func ParseJSONStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

// Story represents
type Story map[string]Chapter

// Chapter represents
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option represents
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
