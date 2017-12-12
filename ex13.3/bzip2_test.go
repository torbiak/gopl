package bzip

import (
	"bufio"
	"bytes"
	"compress/bzip2" // reader
	"fmt"
	"io"
	"strconv"
	"sync"
	"testing"
)

func TestBzip2(t *testing.T) {
	var compressed, uncompressed bytes.Buffer
	w := NewWriter(&compressed)

	// Write a repetitive message in a million pieces,
	// compressing one copy but not the other.
	tee := io.MultiWriter(w, &uncompressed)
	for i := 0; i < 1000000; i++ {
		io.WriteString(tee, "hello")
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	// Check the size of the compressed stream.
	if got, want := compressed.Len(), 255; got != want {
		t.Errorf("1 million hellos compressed to %d bytes, want %d", got, want)
	}

	// Decompress and compare with original.
	var decompressed bytes.Buffer
	io.Copy(&decompressed, bzip2.NewReader(&compressed))
	if !bytes.Equal(uncompressed.Bytes(), decompressed.Bytes()) {
		t.Error("decompression yielded a different message")
	}
}

func TestConcurrentWrites(t *testing.T) {
	n := 100
	c := make(chan int, n)
	for i := 0; i < n; i++ {
		c <- i
	}
	close(c)

	compressed := &bytes.Buffer{}
	w := NewWriter(compressed)
	var err error
	wg := &sync.WaitGroup{}

	consume := func() {
		defer wg.Done()
		for i := range c {
			_, err = w.Write([]byte(fmt.Sprintf("%d\n", i)))
			if err != nil {
				return
			}
		}
	}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go consume()
	}
	wg.Wait()
	if err != nil {
		t.Errorf("%s", err)
	}
	w.Close()

	// Check each number is present.
	seen := make(map[int]bool)
	decompressed := &bytes.Buffer{}
	io.Copy(decompressed, bzip2.NewReader(compressed))
	s := bufio.NewScanner(decompressed)
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			t.Errorf("%s", err)
			return // Corrupted writes?
		}
		seen[i] = true
	}
	var missing []int
	for i := 0; i < n; i++ {
		if !seen[i] {
			missing = append(missing, i)
		}
	}
	if len(missing) > 0 {
		t.Errorf("missing: %s", missing)
	}
}
