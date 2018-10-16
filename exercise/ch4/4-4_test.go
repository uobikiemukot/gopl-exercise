package main

import (
	"testing"
)

func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if (a[i] != b[i]) {
			return false
		}
	}

	return true
}

func TestRotate(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	s1 := make([]int, 6)
	s2 := make([]int, 6)

	for i := 0; i < 12; i++ {
		copy(s1, s)
		copy(s2, s)

		rotateMulti(s1, i)
		rotateSingle(s2, i)

		if !equalSlice(s1, s2) {
			t.Errorf("shift: %d / %v is different from %v\n", i, s2, s1)
		}
	}
}
