package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func wordfreq(r io.Reader) error {
	wc := make(map[string]int)
	total := 0.0

	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		wc[sc.Text()]++
		total++
	}

	for w, c := range wc {
		fmt.Printf("word:%s count:%d freq:%.2f%%\n", w, c, float64(c) / total * 100)
	}

	return nil
}

func main() {
	wordfreq(os.Stdin)
}
