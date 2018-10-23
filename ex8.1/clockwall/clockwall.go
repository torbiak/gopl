// Clockwall connects to a series of clock servers and prints them
// to stdout every second.
//
//Expects a usage statement like this:
// ./wall "New York=8000" Chicago=8010 "Los Angeles=8020"
//
// Output like this:
// New York: 03:04:05	Chicago: 02:04:05	Los Angeles: 00:04:05
//
// Build Note: The folder is called clockwall and so is the program so
// consider a build instruction issued from the ex8.1 root folder like:
// `build -o wall clockwall/clockwall.go`

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type clock struct {
	name       string
	host       string
	connection *net.Conn
	current    string
}

var clocks []*clock

func main() {
	pairs := os.Args[1:]
	for _, item := range pairs {
		tokens := strings.Split(item, "=")
		if len(tokens) == 2 {
			addr := net.JoinHostPort("localhost", tokens[1])
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				log.Fatal(err)
			}
			clocks = append(clocks, &clock{tokens[0], addr, &conn, ""})
		}
	}
	defer closeClocks()

	for {
		for _, c := range clocks {
			s := bufio.NewScanner(*c.connection)
			s.Scan()
			if s.Err() != nil {
				log.Printf("can't read from %s: %s", c.name, s.Err())
			}
			c.current = s.Text()
			fmt.Fprintf(os.Stdout, " %s: %s\t", c.name, c.current)
		}
		fmt.Fprintf(os.Stdout, "\r")
		time.Sleep(100 * time.Millisecond)
	}
}

func closeClocks() {
	for _, c := range clocks {
		(*c.connection).Close()
	}
}

func closeConns(conns []*net.Conn) {
	for _, c := range conns {
		(*c).Close()
	}
}
