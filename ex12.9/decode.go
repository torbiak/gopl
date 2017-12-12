// ex12.9 is a token-based API for decoding s-expressions.
package decode

import (
	"fmt"
	"io"
	"strconv"
	"text/scanner"
)

type Token interface{}
type Symbol string
type String string
type Int int
type StartList struct{}
type EndList struct{}

func (i Int) String() string {
	return fmt.Sprintf("Int(%d)", i)
}

func (s StartList) String() string {
	return "StartList"
}

func (s EndList) String() string {
	return "EndList"
}

type Decoder struct {
	scan  scanner.Scanner
	err   error
	depth int
}

func NewDecoder(r io.Reader) *Decoder {
	var scan scanner.Scanner
	scan.Init(r)
	dec := Decoder{scan: scan}
	scan.Error = dec.setError
	return &dec
}

func (d *Decoder) setError(scan *scanner.Scanner, msg string) {
	d.err = fmt.Errorf("scanning: %s", msg)
}

func (d *Decoder) Token() (Token, error) {
	t := d.scan.Scan()
	if d.err != nil {
		return nil, d.err
	}
	if d.depth == 0 && t != '(' && t != scanner.EOF {
		return nil, fmt.Errorf("expecting '(', got %s", scanner.TokenString(t))
	}
	switch t {
	case scanner.EOF:
		return nil, io.EOF
	case scanner.Ident:
		return Symbol(d.scan.TokenText()), nil
	case scanner.String:
		text := d.scan.TokenText()
		// Assume all strings are quoted.
		return String(text[1 : len(text)-1]), nil
	case scanner.Int:
		n, err := strconv.ParseInt(d.scan.TokenText(), 10, 64)
		if err != nil {
			return nil, err
		}
		return Int(n), nil
	case '(':
		d.depth++
		return StartList{}, nil
	case ')':
		d.depth--
		return EndList{}, nil
	default:
		pos := d.scan.Pos()
		return nil, fmt.Errorf("unexpected token %s at L%d:C%d", scanner.TokenString(t), pos.Line, pos.Column)
	}
}
