package main

import (
	"fmt"
)

func main() {
	var k Kelvin = 0
	fmt.Printf("%v = %v = %v\n", k, KToC(k), KToF(k))
}
