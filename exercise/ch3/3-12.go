package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func reverseString(s string) string {
	var r string

	for i := range s {
		r = string(s[i]) + r
	}

	return string(r)
}

func reverse(s []string) []string {
	var r []string

	for i := range s {
		r = append(r, reverseString(s[i]))
	}

	return r
}

func isReversedPair(a, b string) bool {
	r := reverseString(a)
	return strings.Compare(r, b) == 0
}

func isAnagram2(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	c := []byte(a)
	d := []byte(b)

	for i := range c {
		n := bytes.IndexByte(d, c[i])
		if n == -1 {
			return false
		}
		d[n] = '\x00'
	}

	return true
}

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		n := strings.Index(b, string(a[i]))
		if n == -1 {
			return false
		}
		b = b[0:n] + b[n+1:]
	}

	return true
}

/*
func reverseByte(s string) string {
	b := []byte(s)
	r := make([]byte, 0, len(s))

	for i := range b {
		r = append(b[i:i+1], r)
	}

	return string(r)
}
*/

func main() {
	//fmt.Println(reverse(os.Args[1:]))
	fmt.Println(isAnagram(os.Args[1], os.Args[2]))
}
