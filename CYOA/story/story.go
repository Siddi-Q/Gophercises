package story

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
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
	err := t.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}
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
