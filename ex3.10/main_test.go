package main

import (
	"testing"
)

func TestComma(t *testing.T) {
	tests := []struct {
		s, want string
	}{
		{"1", "1"},
		{"12", "12"},
		{"123", "123"},
		{"1234", "1,234"},
		{"12345", "12,345"},
		{"123456", "123,456"},
		{"1234567", "1,234,567"},
		{"12345678", "12,345,678"},
		{"123456789", "123,456,789"},
		{"1234567890", "1,234,567,890"},
	}
	for _, test := range tests {
		got := comma(test.s)
		if got != test.want {
			t.Errorf("comma(%q), got %q, want %q", test.s, got, test.want)
		}
	}
}
