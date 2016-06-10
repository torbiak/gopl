package shadiff

import (
	"testing"
)

func TestBitDiff(t *testing.T) {
	tests := []struct {
		a, b []byte
		want int
	}{
		{[]byte{0}, []byte{6}, 2},
		{[]byte{1, 2, 3}, []byte{4, 5, 6}, 7},
	}
	for _, test := range tests {
		got := bitDiff(test.a, test.b)
		if got != test.want {
			t.Errorf("bitDiff(%v, %v), got %d, want %d",
				test.a, test.b, got, test.want)
		}
	}
}
