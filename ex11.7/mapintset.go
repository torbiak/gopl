package main

import (
	"bytes"
	"fmt"
	"sort"
)

type MapIntSet struct {
	m map[int]bool
}

func NewMapIntSet() *MapIntSet {
	return &MapIntSet{map[int]bool{}}
}

func (s *MapIntSet) Has(x int) bool {
	return s.m[x]
}

func (s *MapIntSet) Add(x int) {
	s.m[x] = true
}

func (s *MapIntSet) AddAll(nums ...int) {
	for _, x := range nums {
		s.m[x] = true
	}
}

func (s *MapIntSet) UnionWith(t IntSet) {
	for _, x := range t.Ints() {
		s.m[x] = true
	}
}

func (s *MapIntSet) Len() int {
	return len(s.m)
}

func (s *MapIntSet) Remove(x int) {
	delete(s.m, x)
}

func (s *MapIntSet) Clear() {
	s.m = make(map[int]bool)
}

func (s *MapIntSet) Copy() IntSet {
	copy := make(map[int]bool)
	for k, v := range s.m {
		copy[k] = v
	}
	return &MapIntSet{copy}
}

func (s *MapIntSet) String() string {
	b := &bytes.Buffer{}
	b.WriteByte('{')
	for i, x := range s.Ints() {
		if i != 0 {
			b.WriteByte(' ')
		}
		fmt.Fprintf(b, "%d", x)
	}
	b.WriteByte('}')
	return b.String()
}

func (s *MapIntSet) Ints() []int {
	ints := make([]int, 0, len(s.m))
	for x := range s.m {
		ints = append(ints, x)
	}
	sort.IntSlice(ints).Sort()
	return ints
}
