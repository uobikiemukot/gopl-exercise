package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func printMeterFeet(arg string) error {
	l, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		return fmt.Errorf("mf: %v\n", err)
	}

	m := Meter(l)
	f := Feet(l)
	fmt.Printf("%s = %s, %s = %s\n", m, MToF(m), f, FToM(f))

	return nil
}

func main() {
	if len(os.Args) < 2 {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			printMeterFeet(s.Text())
		}
	} else {
		for _, arg := range os.Args[1:] {
			printMeterFeet(arg)
		}
	}
}
