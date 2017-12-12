package main

import (
	"testing"
)

func TestLenZeroInitially(t *testing.T) {
	s := &IntSet{}
	if s.Len() != 0 {
		t.Logf("%d != 0", s.Len())
		t.Fail()
	}
}

func TestLenAfterAddingElements(t *testing.T) {
	s := &IntSet{}
	s.Add(0)
	s.Add(2000)
	if s.Len() != 2 {
		t.Logf("%d != 2", s.Len())
		t.Fail()
	}
}

func TestRemove(t *testing.T) {
	s := &IntSet{}
	s.Add(0)
	s.Remove(0)
	if s.Has(0) {
		t.Log(s)
		t.Fail()
	}
}

func TestClear(t *testing.T) {
	s := &IntSet{}
	s.Add(0)
	s.Add(1000)
	s.Clear()
	if s.Has(0) || s.Has(1000) {
		t.Log(s)
		t.Fail()
	}
}

func TestCopy(t *testing.T) {
	orig := &IntSet{}
	orig.Add(1)
	copy := orig.Copy()
	copy.Add(2)
	if !copy.Has(1) || orig.Has(2) {
		t.Log(orig, copy)
		t.Fail()
	}
}

func TestAddAll(t *testing.T) {
	s := &IntSet{}
	s.AddAll(0, 2, 4)
	if !s.Has(0) || !s.Has(2) || !s.Has(4) {
		t.Log(s)
		t.Fail()
	}
}

func TestIntersectWith(t *testing.T) {
	s := &IntSet{}
	s.AddAll(0, 2, 4)
	u := &IntSet{}
	u.AddAll(1, 2, 3)
	s.IntersectWith(u)
	if !s.Has(2) || s.Len() != 1 {
		t.Log(s)
		t.Fail()
	}
}

func TestDifferenceWith(t *testing.T) {
	s := &IntSet{}
	s.AddAll(0, 2, 4)
	u := &IntSet{}
	u.AddAll(1, 2, 3)
	s.DifferenceWith(u)
	expected := &IntSet{}
	expected.AddAll(0, 4)
	if s.String() != expected.String() {
		t.Log(s)
		t.Fail()
	}
}

func TestSymmetricDifference(t *testing.T) {
	s := &IntSet{}
	s.AddAll(0, 2, 4)
	u := &IntSet{}
	u.AddAll(1, 2, 3)
	s.SymmetricDifference(u)
	expected := &IntSet{}
	expected.AddAll(0, 1, 3, 4)
	if s.String() != expected.String() {
		t.Log(s)
		t.Fail()
	}
}
