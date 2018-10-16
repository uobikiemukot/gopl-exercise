// Rev reverses a slice.
package main

import (
	"fmt"
)

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotateMulti(s []int, n int) {
	if n == 0 {
		return
	}

	if n < 0 {
		n *= -1
		reverse(s)
		reverse(s[:n])
		reverse(s[n:])
	} else {
		reverse(s[:n])
		reverse(s[n:])
		reverse(s)
	}
}

/*
	rotate: 2
	0, 1, 2, 3, 4, 5
	=> 2, 1, 0, 3, 4, 5
	=> 2, 3, 0, 1, 4, 5
	=> 2, 3, 4, 1, 0, 5
	=> 2, 3, 4, 5, 0, 1

	rotate: 4
	0, 1, 2, 3, 4, 5
	=> 4, 1, 2, 3, 0, 5
	=> 4, 5, 2, 3, 0, 1
		rotate: 2
		2, 3, 0, 1
		=> 0, 3, 2, 1
		=> 0, 1, 2, 3
			rotate: 0
	...
	=> 4, 5, 0, 1, 2, 3
*/

func rotateSingle(s []int, n int) {
	if n <= 0 {
		return
	}

	n %= 6
	l := len(s)

	fmt.Printf("s: %v, n: %d, l: %d\n", s, n, l)

	for i := 0; i < l; i++ {
		j := i + n
		if j >= l {
			rotateSingle(s[i:], n - i)
			break
		} else {
			s[i], s[j] = s[j], s[i]
		}
	}
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	t := make([]int, 6)

	for i := 0; i < 7; i++ {
		copy(t, s)
		rotateSingle(t, i)
		fmt.Printf("rotate: %d, s: %v\n", i, t)
	}
}
