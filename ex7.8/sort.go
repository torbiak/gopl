// ex7.8 provides iterative columnar sorting for Persons.
package column

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type columnCmp func(a, b *Person) comparison

type ByColumns struct {
	p       []Person
	columns []columnCmp
}

func NewByColumns(p []Person) *ByColumns {
	return &ByColumns{p, nil}
}

type comparison int

const (
	lt comparison = iota
	eq
	gt
)

func (c *ByColumns) LessName(a, b *Person) comparison {
	switch {
	case a.Name == b.Name:
		return eq
	case a.Name < b.Name:
		return lt
	default:
		return gt
	}
}

func (c *ByColumns) LessAge(a, b *Person) comparison {
	switch {
	case a.Age == b.Age:
		return eq
	case a.Age < b.Age:
		return lt
	default:
		return gt
	}
}

func (c *ByColumns) Len() int      { return len(c.p) }
func (c *ByColumns) Swap(i, j int) { c.p[i], c.p[j] = c.p[j], c.p[i] }

func (c *ByColumns) Less(i, j int) bool {
	for _, f := range c.columns {
		cmp := f(&c.p[i], &c.p[j])
		switch cmp {
		case eq:
			continue
		case lt:
			return true
		case gt:
			return false
		}
	}
	return false
}

func (c *ByColumns) Select(cmp columnCmp) {
	c.columns = append(c.columns, cmp)
	if len(c.columns) > 1 {
		copy(c.columns[1:], c.columns[0:])
		c.columns[0] = cmp
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
