// ex8.8 is a reverb server that disconnects inactive clients.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func scan(r io.Reader, lines chan<- string) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		lines <- s.Text()
	}
	// scan will most likely try to read from the connection after it's closed
	// by handleConn. I don't know how to avoid this. Go seems to shun async io
	// in favour of goroutines, so it probably isn't worth avoiding.
	if s.Err() != nil {
		log.Print("scan: ", s.Err())
	}
}

func handleConn(c net.Conn) {
	wg := &sync.WaitGroup{}
	defer func() {
		wg.Wait()
		c.Close()
	}()
	lines := make(chan string)
	go scan(c, lines)
	timeout := 2 * time.Second
	timer := time.NewTimer(2 * time.Second)
	for {
		select {
		case line := <-lines:
			timer.Reset(timeout)
			wg.Add(1)
			go echo(c, line, 1*time.Second, wg)
		case <-timer.C:
			return
		}
	}
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
