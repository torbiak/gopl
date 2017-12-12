package params

import (
	"testing"
)

func TestPack(t *testing.T) {
	s := struct {
		Name string `http:"n"`
		Age  int    `http:"a"`
	}{"Arugula", 35}
	u, err := Pack(&s)
	if err != nil {
		t.Errorf("Pack(%#v): %s", s, err)
	}
	want := "a=35&n=Arugula"
	got := u.RawQuery
	if got != want {
		t.Errorf("Pack(%#v): got %q, want %q", s, got, want)
	}
}
