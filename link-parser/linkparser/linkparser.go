package linkparser

// Link represents a link/anchor tag
// (<a href="...") in an HTML document.
type Link struct {
	Href string
	Text string
}
