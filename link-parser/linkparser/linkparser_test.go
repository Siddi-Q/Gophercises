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
		{`
<html>
<head>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body>
  <h1>Social stuffs</h1>
  <div>
    <a href="https://www.twitter.com/joncalhoun">
      Check me out on twitter
      <i class="fa fa-twitter" aria-hidden="true"></i>
    </a>
    <a href="https://github.com/gophercises">
      Gophercises is on <strong>Github</strong>!
    </a>
  </div>
</body>
</html>
`, 2, nil, []Link{{Href: "https://www.twitter.com/joncalhoun", Text: "Check me out on twitter"},
			{Href: "https://github.com/gophercises", Text: "Gophercises is on Github!"}}},
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
