package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

const aTag = "a"
const href = "href"

type Link struct {
	Href string
	Text string
}

func main() {
	fileLocationFlag := flag.String(
		"file location",
		"./examples/ex2.html",
		"specify the file location of html to parse",
	)
	flag.Parse()

	// returns a *File which implements the io.Reader interface
	file, err := os.Open(*fileLocationFlag)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		log.Fatalf("could not parse html: %v", err)
	}

	slice := parseHtml(doc, &[]Link{})

	for _, link := range *slice {
		fmt.Printf("%+v", link)
	}

}

func parseHtml(n *html.Node, linkSlice *[]Link) *[]Link {
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
		parseHtml(n.FirstChild, linkSlice)
	}

	if n.NextSibling != nil {
		parseHtml(n.NextSibling, linkSlice)
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
