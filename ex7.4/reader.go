// ex7.4 provides a simple string reader.
package reader

import (
	"io"
)

type stringReader struct {
	s string
}

func (r *stringReader) Read(p []byte) (n int, err error) {
	if len(r.s) == 0 {
		return 0, io.EOF
	}

	n = copy(p, r.s)
	r.s = r.s[n:]

	return
}

func NewReader(s string) io.Reader {
	return &stringReader{s}
}
