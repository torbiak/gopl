package reverse

import (
	"testing"
)

func TestRevUTF8(t *testing.T) {
	s := []byte("Räksmörgås")
	got := string(revUTF8(s))
	want := "sågrömskäR"
	if got != want {
		t.Errorf("got %v, want %v", string(got), want)
	}
}
