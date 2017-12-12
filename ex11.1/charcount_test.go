package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestCharCount(t *testing.T) {
	tests := []struct {
		input   string
		runes   map[rune]int
		props   map[string]int
		sizes   map[int]int
		invalid int
	}{
		{
			input: "Hi, 世.",
			runes: map[rune]int{'H': 1, 'i': 1, ',': 1, ' ': 1, '世': 1, '.': 1},
			props: map[string]int{
				"Ideographic":          1,
				"Pattern_Syntax":       2,
				"Pattern_White_Space":  1,
				"STerm":                1,
				"Sentence_Terminal":    1,
				"Soft_Dotted":          1,
				"Terminal_Punctuation": 2,
				"Unified_Ideograph":    1,
				"White_Space":          1,
			},
			sizes:   map[int]int{1: 5, 3: 1},
			invalid: 0,
		},
	}
	for _, test := range tests {
		runes, props, sizes, invalid := charCount(strings.NewReader(test.input))
		if !reflect.DeepEqual(runes, test.runes) {
			t.Errorf("%q runes: got %v, want %v", test.input, runes, test.runes)
		}
		if !reflect.DeepEqual(props, test.props) {
			t.Errorf("%q props: got %v, want %v", test.input, props, test.props)
		}
		if !reflect.DeepEqual(sizes, test.sizes) {
			t.Errorf("%q sizes: got %v, want %v", test.input, sizes, test.sizes)
		}
		if invalid != test.invalid {
			t.Errorf("%q invalid: got %v, want %v", test.input, invalid, test.invalid)
		}
	}
}
