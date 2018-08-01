package echo

import (
	"fmt"
	"io"
	"strings"
)

func echo1(w io.Writer, args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Fprintln(w, s)
}

func echo2(w io.Writer, args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintln(w, s)
}

func echo3(w io.Writer, args []string) {
	fmt.Fprintln(w, strings.Join(args[1:], " "))
}
