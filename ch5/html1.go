package main

import (
	"fmt"

	"golang.org/x/net/html"
)

var depth int

func forEachNode_(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode_(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.FirstChild == nil {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			for _, attr := range n.Attr {
				fmt.Printf("%*s<%s>\n", depth*2, "", attr.Key)
			}
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		} else {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			for _, attr := range n.Attr {
				fmt.Printf("%*s<%s>\n", depth*2, "", attr.Key)
			}
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
			depth++
		}
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.FirstChild == nil {
			depth--
		}
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
