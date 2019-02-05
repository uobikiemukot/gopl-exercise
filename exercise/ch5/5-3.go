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

// content extract html.Node.Data
// loop version
func content(n *html.Node) []string {
	contents := []string{}

	if n == nil {
		return nil
		//return []string{}
	}

	fmt.Fprintf(os.Stderr, "Type<%v>: %s\n", typeStr[n.Type], n.Data)

	if n.Type == html.TextNode {
		//fmt.Fprintf(os.Stderr, "Type: %v\n", n.Type)
		//fmt.Fprintf(os.Stderr, "Data: %v\n", n.Data)
		//fmt.Fprintf(os.Stderr, "Attr: %v\n", n.Attr)
		contents = append(contents, n.Data)
	}

	if n.Type == html.ElementNode && n.Data == "script" {
		return contents
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		contents = append(contents, content(c)...)
	}

	return contents
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _, s := range content(doc) {
		fmt.Printf("%s", s)
	}
}

