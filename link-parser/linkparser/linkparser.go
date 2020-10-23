package linkparser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a link/anchor tag
// (<a href="...") in an HTML document.
type Link struct {
	Href string
	Text string
}

// Parse will take an HTML document and will return
// a slice of links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	_, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func getLinkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var linkNodes []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		linkNodes = append(linkNodes, getLinkNodes(c)...)
	}
	return linkNodes
}

func buildLink(n *html.Node) Link {
	var link Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			link.Href = attr.Val
			break
		}
	}
	link.Text = getText(n)
	return link
}

func getText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}

	if n.Type != html.ElementNode {
		return ""
	}

	var text string

	for c := n.FirstChild; n != nil; c = n.NextSibling {
		text += getText(c) + " "
	}

	return strings.Join(strings.Fields(text), " ")
}
