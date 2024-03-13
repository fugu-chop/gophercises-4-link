package link

import (
	"strings"

	"golang.org/x/net/html"
)

const aTag = "a"
const href = "href"

type Link struct {
	Href string
	Text string
}

func ParseHtml(n *html.Node, linkSlice *[]Link) *[]Link {
	if n.Type == html.ElementNode && n.Data == aTag {
		for _, a := range n.Attr {
			if a.Key == href {
				newLink := Link{
					Href: a.Val,
					Text: parseLinkText(n),
				}
				*linkSlice = append(*linkSlice, newLink)
			}
		}
	}

	if n.FirstChild != nil {
		ParseHtml(n.FirstChild, linkSlice)
	}

	if n.NextSibling != nil {
		ParseHtml(n.NextSibling, linkSlice)
	}

	return linkSlice
}

func parseLinkText(n *html.Node) string {
	var linkText string

	if n.Type == html.TextNode && len(n.Data) > 0 {
		linkText += n.Data
	}

	if n.FirstChild != nil {
		linkText += parseLinkText(n.FirstChild)
	}

	// The additional logic stops the recursion from digging
	// into neighbouring link tags
	if n.NextSibling != nil && n.NextSibling.Data != aTag {
		linkText += parseLinkText(n.NextSibling)
	}

	return strings.TrimSpace(linkText)
}
