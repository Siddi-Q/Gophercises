package linkparser

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	html := `
<html>
<body>
	<h1>Hello!</h1>
	<a href="/other-page">A link to another page</a>
</body>
</html>
`

	r := strings.NewReader(html)
	links, err := Parse(r)

	if len(links) != 1 || err != nil {
		t.Fatal("Error")
	}

	wantedLink := Link{Href: "/other-page", Text: "A link to another page"}

	if links[0] != wantedLink {
		t.Fatal("Error")
	}
}
