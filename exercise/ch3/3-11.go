// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import (
	"bytes"
	"fmt"
	"os"
)

// comma inserts commas in a non-negative decimal integer string.
func commaSupportFloat(input string) (string, error) {
	ret := new(bytes.Buffer)
	buf := new(bytes.Buffer)
	dec := new(bytes.Buffer)

	// split by '.'
	// initialize buf by
	//   - whole input string (if there is no '.')
	//   - slice[0] of splitting input (if there is one '.')
	//   - else error
	s := bytes.Split([]byte(input), []byte("."))
	if len(s) == 1 {
		buf.Write([]byte(input))
	} else if len(s) == 2 {
		buf.Write(s[0])
		dec.Write(s[1])
	} else {
		return "", fmt.Errorf(`'.' appeared more than once`)
	}

	// check sign (+, -)
	c, _ := buf.ReadByte()
	if c == '-' || c == '+' {
		ret.WriteByte(c)
	} else {
		buf.UnreadByte()
	}

	// initial length
	l := buf.Len()
	t := make([]byte, 3)

	for i := 0; 0 < buf.Len(); i++ {
		if i == 0 && (l%3) != 0 {
			t = buf.Next(l % 3)
		} else {
			t = buf.Next(3)
		}

		if i != 0 {
			ret.WriteString(",")
		}
		ret.Write(t)
	}

	if dec.Len() > 0 {
		ret.WriteString(".")
		ret.Write(dec.Bytes())
	}

	return ret.String(), nil
}

func commaAllowSign(s string) string {
	r := new(bytes.Buffer)
	b := bytes.NewBufferString(s)

	// check sign (+, -)
	c, _ := b.ReadByte()
	if c == '-' || c == '+' {
		r.WriteByte(c)
	} else {
		b.UnreadByte()
	}

	n := b.Len()
	t := make([]byte, 3)

	fmt.Println(b.String())

	for i := 0; 0 < b.Len(); i++ {
		if i == 0 && (n%3) != 0 {
			t = b.Next(n % 3)
		} else {
			t = b.Next(3)
		}

		if i != 0 {
			r.WriteString(",")
		}
		r.Write(t)
	}
	return r.String()
}

func commaNonRecursive(s string) string {
	var t = make([]byte, 3)

	r := new(bytes.Buffer)
	n := len(s)
	b := bytes.NewBufferString(s)

	for i := 0; 0 < b.Len(); i++ {
		if i == 0 && (n%3) != 0 {
			t = b.Next(n % 3)
		} else {
			t = b.Next(3)
		}

		if i != 0 {
			r.WriteString(",")
		}
		r.Write(t)
	}
	return r.String()
}

func commaRecursive(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commaRecursive(s[:n-3]) + "," + s[n-3:]
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		//fmt.Printf("  %s\n", commaNonRecursive(os.Args[i]))
		//fmt.Printf("  %s\n", commaRecursive(os.Args[i]))
		//fmt.Printf("  %s\n", commaAllowSign(os.Args[i]))
		str, _ := commaSupportFloat(os.Args[i])
		fmt.Printf("  %s\n", str)
	}
}
