package main

import (
	"fmt"
	"os"

	"github.com/uobikiemukot/github"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, "usage: edit OWNER REPO ISSUE_NUM\n")
		os.Exit(1)
	}

	ret, err := github.IssueEdit(os.Args[1], os.Args[2], os.Args[3])
	if err != nil {
		fmt.Fprintf(os.Stderr, "github.IssueEdit() failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(ret.Number, ret.HTMLURL)
}
