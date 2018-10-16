package main

import (
	"fmt"
	"os"
	"reflect"
	"unicode"
	"unicode/utf8"
	"unsafe"
)

func squashSpace2(s string) string {
	b := make([]byte, len(s))
	i := 0

	var prev_r rune

	for _, r := range s {
		if !unicode.IsSpace(prev_r) || !unicode.IsSpace(r) {
			i += utf8.EncodeRune(b[i:], r)
		}
		prev_r = r
	}

	return string(b)
}

func squashSpace1(s string) string {
	rs := []rune(s)
	i := 0

	for _, r := range rs {
		//fmt.Printf("index:%d U+%X len:%d\n", i, r, len(string(r)))

		if i > 0 && (unicode.IsSpace(r) && unicode.IsSpace(rs[i-1])) {
			continue
		}

		if unicode.IsSpace(r) {
			rs[i] = 0x20
		} else {
			rs[i] = r
		}
		i++
	}

	/*
		fmt.Println(s)
		fmt.Println(string(rs))
	*/

	return string(rs[:i])
}

func printSliceHeader(s []string) {
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s)))
}

func main() {
	fmt.Println(squashSpace2(os.Args[1]))
}
