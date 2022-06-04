package counter

import (
	"testing"
)

func TestLineCounter(t *testing.T) {
	c := &LineCounter{}
	p := []byte("one\ntwo\nthree\n")
	n, err := c.Write(p)
	if n != len(p) {
		t.Logf("len: %d != %d", n, len(p))
		t.Fail()
	}
	if err != nil {
		t.Log("err: ", err)
		t.Fail()
	}
	if c.N() != 3 {
		t.Logf("lines: %d != 3", c.N())
	}
}

func TestWordCounter(t *testing.T) {
	c := &WordCounter{}
	data := [][]byte{
		[]byte("The upcoming word is sp"),                    // 5
		[]byte("lit across the buffer boundary. "),           // 5
		[]byte(" And this one ends on the buffer boundary."), // 8
		[]byte(" Last words."),                               // 2
	}
	for _, p := range data {
		n, err := c.Write(p)
		if n != len(p) || err != nil {
			t.Logf(`bad write: p="%s" n=%d err="%s"`, string(p), n, err)
			t.Fail()
		}
	}
	if c.N() != 20 {
		t.Logf("words: %d != 20", c.N())
		t.Fail()
	}
}
