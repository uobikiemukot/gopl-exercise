package main

import (
	"bufio"
	"fmt"
	"os"
)

type dup struct {
	count int      /* count of occurence */
	files []string /* file name */
	lines []int    /* line number of file */
}

func main() {
	dups := make(map[string]*dup)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "os.Stdin", dups)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, dups)
			f.Close()
		}
	}
	for line, dup := range dups {
		if dup.count > 1 {
			fmt.Printf("%d\t%s\n", dup.count, line)
			for i := range dup.files {
				fmt.Printf("\t%s:%d\n", dup.files[i], dup.lines[i])
			}
		}
	}
}

func countLines(f *os.File, file string, dups map[string]*dup) {
	input := bufio.NewScanner(f)
	line := 1
	for input.Scan() {
		text := input.Text()
		if _, ok := dups[text]; !ok {
			dups[text] = &dup{}
		}
		dups[text].count++
		dups[text].files = append(dups[text].files, file)
		dups[text].lines = append(dups[text].lines, line)
		line++
	}
	// NOTE: ignoring potential errors from input.Err()
}
