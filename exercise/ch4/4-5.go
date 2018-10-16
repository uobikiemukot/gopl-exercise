package main

import (
	"fmt"
	"os"
)

// remove adjacent duplicate string (in-place)
func uniq(ss []string) []string {
	i := 0
	for _, s := range ss {
		if i == 0 || ss[i-1] != s {
			ss[i] = s
			i++
		}
	}
	return ss[:i]
}

func main() {
	fmt.Println(os.Args[1:])
	fmt.Println(uniq(os.Args[1:]))
}
