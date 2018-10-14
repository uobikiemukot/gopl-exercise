package main

import (
	"fmt"
	"os"

	"github.com/uobikiemukot/github"
)

func main() {
	ret, err := github.IssueSearch(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "github.IssueSearch() failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%d issues:\n", ret.TotalCount)

	for _, i := range ret.Items {
		fmt.Printf("#%-5d %.55s\n", i.Number, i.Title)
	}
}
