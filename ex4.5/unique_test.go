package unique

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	ss := [][]string{
		{"a", "a", "b", "c", "c", "c", "d", "d", "e"},
		{},
	}

	want := [][]string{
		{"a", "b", "c", "d", "e"},
		{},
	}

	for i, s := range ss {
		got := unique(s)
		if !reflect.DeepEqual(got, want[i]) {
			t.Errorf("got %v, want %v", got, want[i])
		}
	}

}
