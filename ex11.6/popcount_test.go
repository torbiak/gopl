// ex11.6 benchmarks popcount implementations.
package popcount

import (
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
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCountTable(x uint64) int {
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

func benchN(b *testing.B, n int, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			f(uint64(j))
		}
	}
}

func benchTableN(b *testing.B, n int) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			PopCountTable(uint64(j))
		}
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

// ClearRightMost should be cheaper than Table (including the table creation)
// if only a few shifts are being performed. Determine how many popcounts make
// Table worth it.
//
// Based on these results the break-even point for Table is slightly over 1000
// invocations versus ClearRightmost and is always cheaper than the shift-based
// methods.
//
// BenchmarkClearRightmost1-4      500000000                4.03 ns/op
// BenchmarkClearRightmost10-4     50000000                37.1 ns/op
// BenchmarkClearRightmost100-4     3000000               538 ns/op
// BenchmarkClearRightmost1000-4     200000              7494 ns/op
// BenchmarkClearRightmost10000-4     20000             92601 ns/op
// BenchmarkTable1-4               200000000                7.66 ns/op
// BenchmarkTable10-4              20000000                79.2 ns/op
// BenchmarkTable100-4              2000000               793 ns/op
// BenchmarkTable1000-4              200000              7975 ns/op
// BenchmarkTable10000-4              20000             79473 ns/op
// BenchmarkShiftValue1-4          20000000                78.3 ns/op
// BenchmarkShiftValue10-4          2000000               782 ns/op
// BenchmarkShiftValue100-4          200000              9203 ns/op
// BenchmarkShiftValue1000-4          10000            119000 ns/op
// BenchmarkShiftValue10000-4          1000           1289813 ns/op
// BenchmarkShiftMask1-4           20000000                78.3 ns/op
// BenchmarkShiftMask10-4           2000000               795 ns/op
// BenchmarkShiftMask100-4           200000              9573 ns/op
// BenchmarkShiftMask1000-4           10000            115398 ns/op
// BenchmarkShiftMask10000-4           1000           1245754 ns/op

func BenchmarkClearRightmost1(b *testing.B) {
	benchN(b, 1, PopCountClearRightmost)
}

func BenchmarkClearRightmost10(b *testing.B) {
	benchN(b, 10, PopCountClearRightmost)
}

func BenchmarkClearRightmost100(b *testing.B) {
	benchN(b, 100, PopCountClearRightmost)
}

func BenchmarkClearRightmost1000(b *testing.B) {
	benchN(b, 1000, PopCountClearRightmost)
}

func BenchmarkClearRightmost10000(b *testing.B) {
	benchN(b, 10000, PopCountClearRightmost)
}

func BenchmarkTable1(b *testing.B) {
	benchTableN(b, 1)
}

func BenchmarkTable10(b *testing.B) {
	benchTableN(b, 10)
}

func BenchmarkTable100(b *testing.B) {
	benchTableN(b, 100)
}

func BenchmarkTable1000(b *testing.B) {
	benchTableN(b, 1000)
}

func BenchmarkTable10000(b *testing.B) {
	benchTableN(b, 10000)
}

func BenchmarkShiftValue1(b *testing.B) {
	benchN(b, 1, PopCountShiftValue)
}

func BenchmarkShiftValue10(b *testing.B) {
	benchN(b, 10, PopCountShiftValue)
}

func BenchmarkShiftValue100(b *testing.B) {
	benchN(b, 100, PopCountShiftValue)
}

func BenchmarkShiftValue1000(b *testing.B) {
	benchN(b, 1000, PopCountShiftValue)
}

func BenchmarkShiftValue10000(b *testing.B) {
	benchN(b, 10000, PopCountShiftValue)
}

func BenchmarkShiftMask1(b *testing.B) {
	benchN(b, 1, PopCountShiftMask)
}

func BenchmarkShiftMask10(b *testing.B) {
	benchN(b, 10, PopCountShiftMask)
}

func BenchmarkShiftMask100(b *testing.B) {
	benchN(b, 100, PopCountShiftMask)
}

func BenchmarkShiftMask1000(b *testing.B) {
	benchN(b, 1000, PopCountShiftMask)
}

func BenchmarkShiftMask10000(b *testing.B) {
	benchN(b, 10000, PopCountShiftMask)
}
