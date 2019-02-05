package main

import (
	"fmt"
	"io"
	"os"
	"strings"

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

func outline(r io.Reader) error {
	doc, err := html.Parse(r)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(*html.Node, bool)) {
	hasChild := (n.FirstChild != nil)

	if pre != nil {
		pre(n, hasChild)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n, hasChild)
	}
}

var depth int

func printTextNode(s string) {
	for _, l := range strings.Split(s, "\n") {
		l = strings.TrimSpace(l)
		if l != "" {
			fmt.Printf("%*s%s\n", depth*2, "", l)
		}
	}
}

func printElementNode(tag string, attr []html.Attribute, hasChild bool) {
	var suffix string

	if hasChild {
		suffix = ">"
	} else {
		suffix = "/>"
	}

	if len(attr) == 0 {
		fmt.Printf("%*s<%s%s\n", depth*2, "", tag, suffix)
		return
	}

	fmt.Printf("%*s<%s ", depth*2, "", tag)

	for i, a := range attr {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%s=\"%s\"", a.Key, a.Val)
	}

	fmt.Println(suffix)
}

func startElement(n *html.Node, hasChild bool) {
	switch n.Type {
	case html.TextNode:
		printTextNode(n.Data)
	case html.DocumentNode:
		// ignore
	case html.ElementNode:
		printElementNode(n.Data, n.Attr, hasChild)
		depth++
	case html.CommentNode:
		fmt.Printf("%*s<!-- %s -->\n", depth*2, "", n.Data)
	case html.DoctypeNode:
		// ignore
	default:
	}
}

func endElement(n *html.Node, hasChild bool) {
	if n.Type == html.ElementNode {
		depth--
		if hasChild {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

func main() {
	// TODO: write test
	err := outline(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
