package decode

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestDecoder(t *testing.T) {
	tests := []struct {
		s            string
		want         []Token
		errSubstring string
	}{
		{`(3 "a" (b))`, []Token{StartList{}, Int(3), String("a"), StartList{}, Symbol("b"), EndList{}, EndList{}}, ""},
		{"(3) a", []Token{StartList{}, Int(3), EndList{}}, "expecting '('"},
		{"(3.2)", []Token{StartList{}}, "unexpected token Float"},
	}
	for _, test := range tests {
		dec := NewDecoder(strings.NewReader(test.s))
		var tokens []Token
		for {
			tok, err := dec.Token()
			if err == io.EOF {
				break
			}
			if err != nil {
				if test.errSubstring != "" {
					if !strings.Contains(err.Error(), test.errSubstring) {
						t.Errorf("decoding %q, expected error containing %s, got %s",
							test.s, test.errSubstring, err)
					}
					break
				} else {
					t.Errorf("decoding %q: %s", test.s, err)
					break
				}
			}
			tokens = append(tokens, tok)
		}
		if !reflect.DeepEqual(tokens, test.want) {
			t.Errorf("Decode(%q), got %s, want %s", test.s, tokens, test.want)
		}
	}
}
