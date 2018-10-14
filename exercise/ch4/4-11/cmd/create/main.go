package main

import (
	"fmt"
	"os"

	"github.com/uobikiemukot/github"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "usage: ic OWNER REPO\n")
		os.Exit(1)
	}

	res, err := github.IssueCreate(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "github.IssueCreate() failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(res.Number)
}
