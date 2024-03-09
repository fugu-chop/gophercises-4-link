package main

import (
	"flag"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	fileLocationFlag := flag.String(
		"file location",
		"./examples/ex1.html",
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

	_ = doc
}
