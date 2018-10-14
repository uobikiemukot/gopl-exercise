package main

import (
	"fmt"
	"os"

	"github.com/uobikiemukot/github"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, "usage: get OWNER REPO ISSUE_NUM\n")
		os.Exit(1)
	}

	res, err := github.IssueGet(os.Args[1], os.Args[2], os.Args[3])
	if err != nil {
		fmt.Fprintf(os.Stderr, "github.IssueGet() failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("title:%s\nbody:%s\n", res.Title, res.Body)
}
