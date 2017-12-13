// ex9.2 lazily initializes the popcount LUT.
package popcount

import (
	"sync"
	"testing"
)

func PopCountShiftMask(x uint64) int {
	count := 0
	mask := uint64(1)
	for i := 0; i < 64; i++ {
		if x&mask > 0 {
			count++
		}
		mask <<= 1
	}
	return count
}

func PopCountShiftValue(x uint64) int {
	count := 0
	mask := uint64(1)
	for i := 0; i < 64; i++ {
		if x&mask > 0 {
			count++
		}
		x >>= 1
	}
	return count
}

func PopCountClearRightmost(x uint64) int {
	count := 0
	for x != 0 {
		x &= x - 1
		count++
	}
	return count
}

// pc[i] is the population count of i.
var loadTableOnce sync.Once
var pc [256]byte

func loadTable() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func Table() [256]byte {
	loadTableOnce.Do(func() {
		for i := range pc {
			pc[i] = pc[i/2] + byte(i&1)
		}
	})
	return pc
}

// PopCount returns the population count (number of set bits) of x.
func PopCountTable(x uint64) int {
	// Calling Table here makes PopCountTable 4 times slower, from 8ns -> 34ns
	// per call.
	pc := Table()
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func bench(b *testing.B, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		f(uint64(i))
	}
}

func BenchmarkTable(b *testing.B) {
	bench(b, PopCountTable)
}

func BenchmarkShiftMask(b *testing.B) {
	bench(b, PopCountShiftMask)
}

func BenchmarkShiftValue(b *testing.B) {
	bench(b, PopCountShiftValue)
}

func BenchmarkClearRightmost(b *testing.B) {
	bench(b, PopCountClearRightmost)
}
