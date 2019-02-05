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

// count count html occurence
// loop version
func count1(n *html.Node, m map[string]int) map[string]int {
	if n == nil {
		return nil
	}

	if n.Type == html.ElementNode {
		m[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		m = count1(c, m)
	}

	return m
}

// count count html occurence
// recursive version
func count2(n *html.Node, m map[string]int) map[string]int {
	if n == nil {
		return m
	}

	if n.Type == html.ElementNode {
		m[n.Data]++
	}

	return merge(count2(n.NextSibling, m), count2(n.FirstChild, m))
}

func merge(a map[string]int, b map[string]int) map[string]int {
	new := map[string]int{}

	for k, v := range a {
		new[k] += v
	}

	for k, v := range b {
		new[k] += v
	}

	return new
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	m := map[string]int{}

	for k, v := range count1(doc, m) {
		fmt.Println(k, v)
	}
}

