package main

import (
	"math/rand"
	"testing"
)

func newIntSets() []IntSet {
	return []IntSet{&BitIntSet{}, &BitIntSet32{}, NewMapIntSet()}
}

func TestLenZeroInitially(t *testing.T) {
	for _, s := range newIntSets() {
		if s.Len() != 0 {
			t.Errorf("%T.Len(): got %d, want 0", s, s.Len())
		}
	}
}

func TestLenAfterAddingElements(t *testing.T) {
	for _, s := range newIntSets() {
		s.Add(0)
		s.Add(2000)
		if s.Len() != 2 {
			t.Errorf("%T.Len(): got %d, want 2", s, s.Len())
		}
	}
}

func TestRemove(t *testing.T) {
	for _, s := range newIntSets() {
		s.Add(0)
		s.Remove(0)
		if s.Has(0) {
			t.Errorf("%T: want zero removed, got %s", s, s)
		}
	}
}

func TestClear(t *testing.T) {
	for _, s := range newIntSets() {
		s.Add(0)
		s.Add(1000)
		s.Clear()
		if s.Has(0) || s.Has(1000) {
			t.Errorf("%T: want empty set, got %s", s, s)
		}
	}
}

func TestCopy(t *testing.T) {
	for _, orig := range newIntSets() {
		orig.Add(1)
		copy := orig.Copy()
		copy.Add(2)
		if !copy.Has(1) || orig.Has(2) {
			t.Errorf("%T: want %s, got %s", orig, orig, copy)
		}
	}
}

func TestAddAll(t *testing.T) {
	for _, s := range newIntSets() {
		s.AddAll(0, 2, 4)
		if !s.Has(0) || !s.Has(2) || !s.Has(4) {
			t.Errorf("%T: want {2 4}, got %s", s, s)
		}
	}
}

const max = 32000

func addRandom(set IntSet, n int) {
	for i := 0; i < n; i++ {
		set.Add(rand.Intn(max))
	}
}

func benchHas(b *testing.B, set IntSet, n int) {
	addRandom(set, n)
	for i := 0; i < b.N; i++ {
		set.Has(rand.Intn(max))
	}
}

func benchAdd(b *testing.B, set IntSet, n int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			set.Add(rand.Intn(max))
		}
		set.Clear()
	}
}

func randInts(n int) []int {
	ints := make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = rand.Intn(max)
	}
	return ints
}

func benchAddAll(b *testing.B, set IntSet, batchSize int) {
	ints := randInts(batchSize)
	for i := 0; i < b.N; i++ {
		set.AddAll(ints...)
		set.Clear()
	}
}

func benchUnionWith(bm *testing.B, a, b IntSet, n int) {
	addRandom(a, n)
	addRandom(b, n)
	for i := 0; i < bm.N; i++ {
		a.UnionWith(b)
	}
}

func benchString(b *testing.B, set IntSet, n int) {
	addRandom(set, n)
	for i := 0; i < b.N; i++ {
		set.String()
	}
}

//func Benchmark<Type><Method><Size>(b *testing.B) {
//	bench<Method>(b, New<Type>(), <Size>)
//}
func BenchmarkMapIntSetAdd10(b *testing.B) {
	benchAdd(b, NewMapIntSet(), 10)
}
func BenchmarkMapIntSetAdd100(b *testing.B) {
	benchAdd(b, NewMapIntSet(), 100)
}
func BenchmarkMapIntSetAdd1000(b *testing.B) {
	benchAdd(b, NewMapIntSet(), 1000)
}
func BenchmarkMapIntSetHas10(b *testing.B) {
	benchHas(b, NewMapIntSet(), 10)
}
func BenchmarkMapIntSetHas100(b *testing.B) {
	benchHas(b, NewMapIntSet(), 100)
}
func BenchmarkMapIntSetHas1000(b *testing.B) {
	benchHas(b, NewMapIntSet(), 1000)
}
func BenchmarkMapIntSetAddAll10(b *testing.B) {
	benchAddAll(b, NewMapIntSet(), 10)
}
func BenchmarkMapIntSetAddAll100(b *testing.B) {
	benchAddAll(b, NewMapIntSet(), 100)
}
func BenchmarkMapIntSetAddAll1000(b *testing.B) {
	benchAddAll(b, NewMapIntSet(), 1000)
}
func BenchmarkMapIntSetString10(b *testing.B) {
	benchString(b, NewMapIntSet(), 10)
}
func BenchmarkMapIntSetString100(b *testing.B) {
	benchString(b, NewMapIntSet(), 100)
}
func BenchmarkMapIntSetString1000(b *testing.B) {
	benchString(b, NewMapIntSet(), 1000)
}
func BenchmarkBitIntSetAdd10(b *testing.B) {
	benchAdd(b, NewBitIntSet(), 10)
}
func BenchmarkBitIntSetAdd100(b *testing.B) {
	benchAdd(b, NewBitIntSet(), 100)
}
func BenchmarkBitIntSetAdd1000(b *testing.B) {
	benchAdd(b, NewBitIntSet(), 1000)
}
func BenchmarkBitIntSetHas10(b *testing.B) {
	benchHas(b, NewBitIntSet(), 10)
}
func BenchmarkBitIntSetHas100(b *testing.B) {
	benchHas(b, NewBitIntSet(), 100)
}
func BenchmarkBitIntSetHas1000(b *testing.B) {
	benchHas(b, NewBitIntSet(), 1000)
}
func BenchmarkBitIntSetAddAll10(b *testing.B) {
	benchAddAll(b, NewBitIntSet(), 10)
}
func BenchmarkBitIntSetAddAll100(b *testing.B) {
	benchAddAll(b, NewBitIntSet(), 100)
}
func BenchmarkBitIntSetAddAll1000(b *testing.B) {
	benchAddAll(b, NewBitIntSet(), 1000)
}
func BenchmarkBitIntSetString10(b *testing.B) {
	benchString(b, NewBitIntSet(), 10)
}
func BenchmarkBitIntSetString100(b *testing.B) {
	benchString(b, NewBitIntSet(), 100)
}
func BenchmarkBitIntSetString1000(b *testing.B) {
	benchString(b, NewBitIntSet(), 1000)
}
func BenchmarkBitIntSet32Add10(b *testing.B) {
	benchAdd(b, NewBitIntSet32(), 10)
}
func BenchmarkBitIntSet32Add100(b *testing.B) {
	benchAdd(b, NewBitIntSet32(), 100)
}
func BenchmarkBitIntSet32Add1000(b *testing.B) {
	benchAdd(b, NewBitIntSet32(), 1000)
}
func BenchmarkBitIntSet32Has10(b *testing.B) {
	benchHas(b, NewBitIntSet32(), 10)
}
func BenchmarkBitIntSet32Has100(b *testing.B) {
	benchHas(b, NewBitIntSet32(), 100)
}
func BenchmarkBitIntSet32Has1000(b *testing.B) {
	benchHas(b, NewBitIntSet32(), 1000)
}
func BenchmarkBitIntSet32AddAll10(b *testing.B) {
	benchAddAll(b, NewBitIntSet32(), 10)
}
func BenchmarkBitIntSet32AddAll100(b *testing.B) {
	benchAddAll(b, NewBitIntSet32(), 100)
}
func BenchmarkBitIntSet32AddAll1000(b *testing.B) {
	benchAddAll(b, NewBitIntSet32(), 1000)
}
func BenchmarkBitIntSet32String10(b *testing.B) {
	benchString(b, NewBitIntSet32(), 10)
}
func BenchmarkBitIntSet32String100(b *testing.B) {
	benchString(b, NewBitIntSet32(), 100)
}
func BenchmarkBitIntSet32String1000(b *testing.B) {
	benchString(b, NewBitIntSet32(), 1000)
}

//func BenchMark<Type>UnionWith<Size>(b *testing.B) {
//	benchUnionWith(b, New<Type>(), New<Type>(), <Size>)
//}
func BenchMarkMapIntSetUnionWith10(b *testing.B) {
	benchUnionWith(b, NewMapIntSet(), NewMapIntSet(), 10)
}
func BenchMarkMapIntSetUnionWith100(b *testing.B) {
	benchUnionWith(b, NewMapIntSet(), NewMapIntSet(), 100)
}
func BenchMarkMapIntSetUnionWith1000(b *testing.B) {
	benchUnionWith(b, NewMapIntSet(), NewMapIntSet(), 1000)
}
func BenchMarkBitIntSetUnionWith10(b *testing.B) {
	benchUnionWith(b, NewBitIntSet(), NewBitIntSet(), 10)
}
func BenchMarkBitIntSetUnionWith100(b *testing.B) {
	benchUnionWith(b, NewBitIntSet(), NewBitIntSet(), 100)
}
func BenchMarkBitIntSetUnionWith1000(b *testing.B) {
	benchUnionWith(b, NewBitIntSet(), NewBitIntSet(), 1000)
}
func BenchMarkBitIntSet32UnionWith10(b *testing.B) {
	benchUnionWith(b, NewBitIntSet32(), NewBitIntSet32(), 10)
}
func BenchMarkBitIntSet32UnionWith100(b *testing.B) {
	benchUnionWith(b, NewBitIntSet32(), NewBitIntSet32(), 100)
}
func BenchMarkBitIntSet32UnionWith1000(b *testing.B) {
	benchUnionWith(b, NewBitIntSet32(), NewBitIntSet32(), 1000)
}

// BenchmarkMapIntSetAdd10-4                 500000              2426 ns/op             323 B/op           3 allocs/op
// BenchmarkMapIntSetAdd100-4                 50000             35449 ns/op            3471 B/op          20 allocs/op
// BenchmarkMapIntSetAdd1000-4                 5000            439968 ns/op           55342 B/op          98 allocs/op
// BenchmarkMapIntSetHas10-4               20000000                82.0 ns/op             0 B/op           0 allocs/op
// BenchmarkMapIntSetHas100-4              20000000                85.3 ns/op             0 B/op           0 allocs/op
// BenchmarkMapIntSetHas1000-4             20000000                87.7 ns/op             0 B/op           0 allocs/op
// BenchmarkMapIntSetAddAll10-4             1000000              1701 ns/op             323 B/op           3 allocs/op
// BenchmarkMapIntSetAddAll100-4             100000             18160 ns/op            3475 B/op          20 allocs/op
// BenchmarkMapIntSetAddAll1000-4              5000            227188 ns/op           55360 B/op          98 allocs/op
// BenchmarkMapIntSetString10-4              300000              4550 ns/op             368 B/op          14 allocs/op
// BenchmarkMapIntSetString100-4              30000             50945 ns/op            4569 B/op         107 allocs/op
// BenchmarkMapIntSetString1000-4              3000            512665 ns/op           40205 B/op        1001 allocs/op
// BenchmarkBitIntSetAdd10-4                2000000               703 ns/op               0 B/op           0 allocs/op
// BenchmarkBitIntSetAdd100-4                200000              5933 ns/op               0 B/op           0 allocs/op
// BenchmarkBitIntSetAdd1000-4                20000             57948 ns/op               0 B/op           0 allocs/op
// BenchmarkBitIntSetHas10-4               30000000                56.2 ns/op             0 B/op           0 allocs/op
// BenchmarkBitIntSetHas100-4              30000000                53.6 ns/op             0 B/op           0 allocs/op
// BenchmarkBitIntSetHas1000-4             30000000                53.4 ns/op             0 B/op           0 allocs/op
// BenchmarkBitIntSetAddAll10-4            10000000               181 ns/op               0 B/op           0 allocs/op
// BenchmarkBitIntSetAddAll100-4            2000000               947 ns/op               0 B/op           0 allocs/op
// BenchmarkBitIntSetAddAll1000-4            200000              8287 ns/op               0 B/op           0 allocs/op
// BenchmarkBitIntSetString10-4              300000              5966 ns/op             256 B/op          12 allocs/op
// BenchmarkBitIntSetString100-4              30000             54357 ns/op            3649 B/op         106 allocs/op
// BenchmarkBitIntSetString1000-4              5000            394976 ns/op           31954 B/op         998 allocs/op
// BenchmarkBitIntSet32Add10-4              2000000               669 ns/op               0 B/op           0 allocs/op
// BenchmarkBitIntSet32Add100-4              300000              5774 ns/op               0 B/op           0 allocs/op
// BenchmarkBitIntSet32Add1000-4              30000             55189 ns/op               0 B/op           0 allocs/op
// BenchmarkBitIntSet32Has10-4             30000000                53.9 ns/op             0 B/op           0 allocs/op
// BenchmarkBitIntSet32Has100-4            30000000                52.2 ns/op             0 B/op           0 allocs/op
// BenchmarkBitIntSet32Has1000-4           30000000                57.2 ns/op             0 B/op           0 allocs/op
// BenchmarkBitIntSet32AddAll10-4          10000000               186 ns/op               0 B/op           0 allocs/op
// BenchmarkBitIntSet32AddAll100-4          2000000               965 ns/op               0 B/op           0 allocs/op
// BenchmarkBitIntSet32AddAll1000-4          200000              9187 ns/op               0 B/op           0 allocs/op
// BenchmarkBitIntSet32String10-4            300000              5230 ns/op             256 B/op          12 allocs/op
// BenchmarkBitIntSet32String100-4            30000             47954 ns/op            3649 B/op         106 allocs/op
// BenchmarkBitIntSet32String1000-4            3000            410110 ns/op           31898 B/op         991 allocs/op
