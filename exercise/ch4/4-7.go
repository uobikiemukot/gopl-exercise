package main

import (
	"fmt"
	"os"
)

func reverse(s string) string {
	rs := []rune(s)
	l := len(rs)

	for i := 0; i < l/2; i++ {
		rs[i], rs[(l-1)-i] = rs[(l-1)-i], rs[i]
	}

	return string(rs)
}

func main() {
	fmt.Println(reverse(os.Args[1]))
}
