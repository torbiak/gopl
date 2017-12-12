package unique

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	s := []string{"a", "a", "b", "c", "c", "c", "d", "d", "e"}
	got := unique(s)
	want := []string{"a", "b", "c", "d", "e"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
