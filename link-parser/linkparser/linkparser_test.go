package linkparser

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	var tests = []struct {
		html   string
		length int
		err    error
		links  []Link
	}{
		{`
<html>
<body>
	<h1>Hello!</h1>
	<a href="/other-page">A link to another page</a>
</body>
</html>
`, 1, nil, []Link{{Href: "/other-page", Text: "A link to another page"}}},
	}

	for _, test := range tests {
		r := strings.NewReader(test.html)
		links, err := Parse(r)

		if len(links) != test.length || err != test.err {
			t.Fatalf(`Got %d, %v. Want %d, %v.`, len(links), err, test.length, test.err)
		}

		for i, link := range test.links {
			if links[i] != link {
				t.Fatalf(`Got %+v. Want %+v.`, links[i], link)
			}
		}
	}
}
