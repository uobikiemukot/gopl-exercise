// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

const (
	Control = iota
	Letter
	Mark
	Number
	Punct
	Space
)

var Category = []string{
	"Ctrl",
	"Letter",
	"Mark",
	"Number",
	"Punct",
	"Space",
}

func charcount(r io.Reader) error {
	var cc [Space + 1]int

	in := bufio.NewReader(r)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			return err
		}
		if r == unicode.ReplacementChar && n == 1 {
			continue
		}

		if unicode.IsControl(r) {
			cc[Control]++
		}
		if unicode.IsLetter(r) {
			cc[Letter]++
		}
		if unicode.IsMark(r) {
			cc[Mark]++
		}
		if unicode.IsNumber(r) {
			cc[Number]++
		}
		if unicode.IsPunct(r) {
			cc[Punct]++
		}
		if unicode.IsSpace(r) {
			cc[Space]++
		}
	}

	for i, c := range cc {
		fmt.Printf("category:%s\tcount:%d\n", Category[i], c)
	}

	return nil
}

func main() {
	if err := charcount(os.Stdin); err != nil {
		os.Exit(1)
	}
}
