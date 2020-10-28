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
		t.Fatalf(`Got %d, %v. Want %d, %v.`, len(links), err, 1, nil)
	}

	wantedLink := Link{Href: "/other-page", Text: "A link to another page"}

	if links[0] != wantedLink {
		t.Fatalf(`Got %+v. Want %+v.`, links[0], wantedLink)
	}
}
