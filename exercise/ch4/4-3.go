// Rev reverses a slice.
package main

import (
	"fmt"
)

const (
	Size = 6 // array size
)

// reverse reverses a slice of ints in place.
func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseArray(a *[Size]int) {
	for i, j := 0, Size-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	a := [Size]int{0, 1, 2, 3, 4, 5}
	reverseArray(&a)
	//reverseArray(a)
	fmt.Println(a) // "[5 4 3 2 1 0]"

	/*
		s := []int{0, 1, 2, 3, 4, 5}
		// Rotate s left by two positions.
		reverse(s[:2])
		reverse(s[2:])
		reverse(s)
		fmt.Println(s) // "[2 3 4 5 0 1]"
	*/

	// Interactive test of reverse.
	/*
			input := bufio.NewScanner(os.Stdin)
		outer:
			for input.Scan() {
				var ints []int
				for _, s := range strings.Fields(input.Text()) {
					x, err := strconv.ParseInt(s, 10, 64)
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
						continue outer
					}
					ints = append(ints, int(x))
				}
				reverse(ints)
				fmt.Printf("%v\n", ints)
			}
	*/
}
