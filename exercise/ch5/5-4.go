package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

/*
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
*/

var typeStr = map[html.NodeType]string{
	html.ErrorNode:    "ErrorNote",
	html.TextNode:     "TextNode",
	html.DocumentNode: "DocumentNote",
	html.ElementNode:  "ElementNode",
	html.CommentNode:  "CommentNode",
	html.DoctypeNode:  "DoctypeNode",
}

func attr(q string, attr []html.Attribute) string {
	for _, a := range attr {
		if a.Key == q {
			return a.Val
		}
	}
	return ""
}

// link extract various links (a, img, script, link)
func link(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			links = append(links, "a:href:" + attr("href", n.Attr))
		case "link":
			links = append(links, "link:href:" + attr("href", n.Attr))
		case "img":
			links = append(links, "img:src:" + attr("src", n.Attr))
		case "script":
			links = append(links, "script:src:" + attr("src", n.Attr))
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = link(links, c)
	}

	return links
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _, s := range link([]string{}, doc) {
		fmt.Printf("%s\n", s)
	}
}

