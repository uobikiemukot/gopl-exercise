package main

import (
	"fmt"
	"os"
	"strings"
)

func expand(s string, f func(string) string) string {
	l := make([]string, 0)


	for _, s := range strings.Split(s, " ") {
		if strings.HasPrefix(s, "$") {
			l = append(l, f(s[1:]))
		} else {
			l = append(l, s)
		}
	}

	return strings.Join(l, " ")
}

func main() {
	for _, s := range os.Args[1:] {
		fmt.Println(expand(s, strings.ToUpper))
	}
}
