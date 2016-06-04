package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	b := bytes.Buffer{}
	mantissaStart := 0
	if s[0] == '+' || s[0] == '-' {
		b.WriteByte(s[0])
		mantissaStart = 1
	}
	mantissaEnd := strings.Index(s, ".")
	if mantissaEnd == -1 {
		mantissaEnd = len(s)
	}
	mantissa := s[mantissaStart:mantissaEnd]
	pre := len(mantissa) % 3
	if pre > 0 {
		b.Write([]byte(mantissa[:pre]))
		if len(mantissa) > pre {
			b.WriteString(",")
		}
	}
	for i, c := range mantissa[pre:] {
		if i%3 == 0 && i != 0 {
			b.WriteString(",")
		}
		b.WriteRune(c)
	}
	b.WriteString(s[mantissaEnd:])
	return b.String()
}

func main() {
	fmt.Println(comma("1"))
	fmt.Println(comma("12"))
	fmt.Println(comma("123"))
	fmt.Println(comma("1234"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567"))
	fmt.Println(comma("12345678"))
	fmt.Println(comma("123456789"))
	fmt.Println(comma("1234567890"))

	fmt.Println(comma("1.1234"))
	fmt.Println(comma("12.1234"))
	fmt.Println(comma("123.1234"))
	fmt.Println(comma("1234.1234"))
	fmt.Println(comma("12345.1234"))
	fmt.Println(comma("123456.1234"))
	fmt.Println(comma("1234567.1234"))
	fmt.Println(comma("12345678.1234"))
	fmt.Println(comma("123456789.1234"))
	fmt.Println(comma("1234567890.1234"))

	fmt.Println(comma("+123456789.1234"))
	fmt.Println(comma("-1234567890.1234"))
}
