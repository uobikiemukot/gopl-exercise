// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"testing"
)

// Test

func testPopCount(t *testing.T, popCount func(uint64) int, m map[uint64]int) {
	for k, v := range m {
		want := v
		got := popCount(k)
		if want != got {
			t.Errorf("value: %d want: %d got: %d\n", k, want, got)
		}
	}
}

func TestPopCount(t *testing.T) {
	m := map[uint64]int{
		0: 0, 1: 1, 2: 1, 3: 2, 4: 1, 5: 2, 6: 2, 7: 3,
		8: 1, 9: 2, 10: 2, 11: 3, 12: 2, 13: 3, 14: 3, 15: 4,
	}

	t.Run("TableLookUp", func(t *testing.T) { testPopCount(t, PopCountByTableLookUp, m) })
	t.Run("TableLookUp2", func(t *testing.T) { testPopCount(t, PopCountByTableLookUp2, m) })
	t.Run("Shift", func(t *testing.T) { testPopCount(t, PopCountByShift, m) })
	t.Run("Clear", func(t *testing.T) { testPopCount(t, PopCountByClear, m) })
}

// BenchMark

func benchPopCount(b *testing.B, popCount func(uint64) int, v uint64) {
	for i := 0; i < b.N; i++ {
		popCount(v)
	}
}

func BenchmarkPopCount(b *testing.B) {
	var v uint64 = 0x1234567890ABCDEF

	b.Run("TableLookUp", func(b *testing.B) { benchPopCount(b, PopCountByTableLookUp, v) })
	b.Run("TableLookUp2", func(b *testing.B) { benchPopCount(b, PopCountByTableLookUp2, v) })
	b.Run("Shift", func(b *testing.B) { benchPopCount(b, PopCountByShift, v) })
	b.Run("Clear", func(b *testing.B) { benchPopCount(b, PopCountByClear, v) })
}
