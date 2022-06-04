// ex7.1 provides line and word counters.
package counter

import (
	"bufio"
	"fmt"
)

type LineCounter struct {
	lines int
}

func (c *LineCounter) Write(p []byte) (n int, err error) {
	for i, advance, atEOF := 0, 0, false; i < len(p); i += advance {
		var token []byte
		advance, token, _ = bufio.ScanLines(p[i:], atEOF)
		if token == nil {
			atEOF = true
			continue
		}
		c.lines++
	}
	return len(p), nil
}

func (c *LineCounter) N() int {
	return c.lines
}

func (c *LineCounter) String() string {
	return fmt.Sprintf("%d", c.lines)
}

type WordCounter struct {
	words  int
	inWord bool
}

func (c *WordCounter) Write(p []byte) (n int, err error) {
	for i, advance, atEOF := 0, 0, false; i < len(p); i += advance {
		var token []byte
		advance, token, _ = bufio.ScanWords(p[i:], atEOF)
		// according to the source code of bufio.ScanWords,
		// we should request the final, incomplete word explicitly with atEOF set to true
		if token == nil {
			atEOF = true
			continue
		}
		c.words++
	}
	return len(p), nil
}

func (c *WordCounter) N() int {
	return c.words
}

func (c *WordCounter) String() string {
	return fmt.Sprintf("%d", c.words)
}
