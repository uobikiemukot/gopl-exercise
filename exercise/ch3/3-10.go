// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import (
	"bytes"
	"fmt"
	"os"
)

const (
	split = 3
)

// comma inserts commas in a non-negative decimal integer string.
func commaNonRecursive(s string) string {
	/*
		b := []byte(s)
		n := bytes.SplitN(b, nil, len(s)/3)
		return string(bytes.Join(n, []byte(",")))
	*/
	//var res string
	var t = make([]byte, 3)

	r := new(bytes.Buffer)
	n := len(s)
	b := bytes.NewBufferString(s)

	for i := 0; 0 < b.Len(); i++ {
		if i == 0 && (n % 3) != 0 {
			t = b.Next(n % 3)
		} else {
			t = b.Next(3)
		}

		/*
		if i != 0 {
			res += ","
		}
		res += string(t)
		return res
		*/
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
		fmt.Printf("  %s\n", commaNonRecursive(os.Args[i]))
		//fmt.Printf("  %s\n", commaRecursive(os.Args[i]))
	}
}
