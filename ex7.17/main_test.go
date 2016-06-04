package main

import (
	"testing"
	"bytes"
	"strings"
)

func TestParseSelectors(t *testing.T) {
	sels, err := parseSelectors(`a p [href="hi"][id=big] p[class="blue"] [class]`)
	if err != nil {
		t.Error(err)
	}
	expected := []string{
		"a", "p", `[href="hi"][id="big"]`, `p[class="blue"]`, "[class]",
	}
	if len(expected) != len(sels) {
		t.Errorf("%s != %s", sels, expected)
		return
	}
	for i, ex := range expected {
		actual := sels[i].String()
		if actual != ex {
			t.Errorf("%s != %s and %s != %s", actual, ex, sels, expected)
			return
		}
	}
}

func TestSelectorParseFailure_badAttr(t *testing.T) {
	sels, err := parseSelectors("a]")
	if !strings.Contains(err.Error(), "want ident") {
		t.Error(sels, err)
	}
}

func TestSelectorParseFailure_badTag(t *testing.T) {
	sels, err := parseSelectors(`a "p"`)
	if !strings.Contains(err.Error(), "want ident") {
		t.Error(sels, err)
	}
}

func TestXMLSelect(t *testing.T) {
	tests := []struct {
		selectors, xml string
	}{
		{`a[id="3"] [id="4"]`, `<a id="3"><p id="4">good</p></a><a><p id="4">bad</p></a>`},
		{`[id="3"] [id]`, `<a id="3"><p id="4">good</p></a><a><p id="4">bad</p></a>`},
		{`[id="3"][class=big]`, `<a id="3" class="big">good</a><a id="3">bad</a>`},
	}
	for _, test := range tests {
		r := strings.NewReader(test.xml)
		w := &bytes.Buffer{}
		sels, err := parseSelectors(test.selectors)
		if err != nil {
			t.Error(test, err)
			return
		}
		xmlselect(w, r, sels)
		if w.String() != "good\n" {
			t.Errorf("%s: %q != \"good\\n\"", test, w.String())
		}
	}
}
