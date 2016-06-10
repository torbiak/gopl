// ex11.7 benchmarks IntSet implementations.
package main

type IntSet interface {
	Has(x int) bool
	Add(x int)
	AddAll(nums ...int)
	UnionWith(t IntSet)
	Len() int
	Remove(x int)
	Clear()
	Copy() IntSet
	String() string
	Ints() []int
}
