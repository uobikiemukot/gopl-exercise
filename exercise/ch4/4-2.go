package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func popCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func countDiffBits(x, y [sha256.Size]byte) int {
	sum := 0
	for i := 0; i < sha256.Size; i++ {
		sum += popCount(uint64(x[i] ^ y[i]))
	}
	return sum
}

func main() {
	use384 := flag.Bool("384", false, "use sha512.Sum384")
	use512 := flag.Bool("512", false, "use sha512.Sum512")
	flag.Parse()

	b := new(bytes.Buffer)
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		b.Write(s.Bytes())
	}

	if *use384 {
		r := sha512.Sum384(b.Bytes())
		fmt.Printf("%x\n", r)
	} else if *use512 {
		r := sha512.Sum512(b.Bytes())
		fmt.Printf("%x\n", r)
	} else {
		r := sha256.Sum256(b.Bytes())
		fmt.Printf("%x\n", r)
	}
}
