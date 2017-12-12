package cyclic

import (
	"testing"
)

type node struct {
	next *node
}

func TestCyclic(t *testing.T) {
	a, b, c := node{}, node{}, node{}
	a.next = &b
	b.next = &c
	acyclicList := a
	d, e, f := node{}, node{}, node{}
	d.next = &e
	e.next = &f
	f.next = &d
	cyclicList := d

	tests := []struct {
		x    interface{}
		want bool
	}{
		{1, false},
		{nil, false},
		{cyclicList, true},
		{acyclicList, false},
		{[]node{cyclicList}, true},
		{[]node{acyclicList}, false},
		{map[node]int{acyclicList: 1}, false},
		{map[node]int{cyclicList: 1}, true},
		{map[int]node{1: acyclicList}, false},
		{map[int]node{1: cyclicList}, true},
		{&cyclicList, true},
	}
	for _, test := range tests {
		if Cyclic(test.x) != test.want {
			t.Errorf("Cyclic(%v) != %v", test.x, !test.want)
		}
	}
}
