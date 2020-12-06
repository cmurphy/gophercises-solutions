package link

import (
	"golang.org/x/net/html"
	"io"
	"strings"
)

// Link represents the href and text of an HTML <a> element.
type Link struct {
	Href string
	Text string
}

// Parse accepts an HTML document and returns a slice of links found in it or an error.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	return scanTree(doc), nil
}

func scanTree(n *html.Node) []Link {
	links := make([]Link, 0)
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			text := getLinkText(n)
			if a.Key == "href" {
				links = append(links, Link{Href: a.Val, Text: text})
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, scanTree(c)...)
	}
	return links
}

func getLinkText(n *html.Node) string {
	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			text = text + c.Data
		}
		if c.Type == html.ElementNode {
			text = text + getLinkText(c)
		}
	}
	return strings.TrimSpace(text)
}
