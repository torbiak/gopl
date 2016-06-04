package main

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	xml := `<doc><a id="b"><b/>hi<b>rah</b></a></doc>`
	node, err := parse(strings.NewReader(xml))
	if err != nil {
		t.Error(err)
	}
	el := node.(*Element)
	expected := `doc []
  a [{{ id} b}]
    b []
    "hi"
    b []
      "rah"
`
	if el.String() != expected {
		t.Errorf("%q != %q")
	}
}
