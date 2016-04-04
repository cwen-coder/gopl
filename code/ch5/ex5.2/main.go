package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "countNode: %v\n", err)
		os.Exit(1)
	}
	dataMap := make(map[string]int)
	for k, v := range countNode(dataMap, doc) {
		fmt.Printf("%v : %v\n", k, v)
	}
}

func countNode(data map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		data[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		data = countNode(data, c)
	}
	return data
}
