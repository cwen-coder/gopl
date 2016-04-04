package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "textVal: %v\n", err)
		os.Exit(1)
	}
	textVal(doc)
}
func textVal(n *html.Node) {
	if n.Type == html.TextNode {
		if s := strings.TrimSpace(n.Data); len(s) > 0 {
			fmt.Printf("%s\n", s)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data != "script" && c.Data != "stype" {
			textVal(c)
		}
	}
}
