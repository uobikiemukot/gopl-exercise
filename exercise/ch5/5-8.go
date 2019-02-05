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

func forEachNode(n *html.Node, id string, pre, post func(*html.Node, string) bool) *html.Node {
	if pre != nil {
		if pre(n, id) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if n := forEachNode(c, id, pre, post); n != nil {
			return n
		}
	}

	if post != nil {
		if post(n, id) {
			return n
		}
	}

	return nil
}

func preMatch(n *html.Node, id string) bool {
	if n.Type != html.ElementNode {
		return false
	}
	fmt.Fprintln(os.Stderr, "pre:" + n.Data)

	return n.Data == id
}

func postMatch(n *html.Node, id string) bool {
	if n.Type != html.ElementNode {
		return false
	}
	fmt.Fprintln(os.Stderr, "post:" + n.Data)

	return n.Data == id
}

func ElementByID(doc *html.Node, id string) *html.Node {
	n := forEachNode(doc, id, preMatch, postMatch)
	return n
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		os.Exit(1)
	}

	n := ElementByID(doc, os.Args[1]) 
	fmt.Println(n)
}
