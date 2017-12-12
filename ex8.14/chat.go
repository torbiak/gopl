// ex8.14 is a chat server that prompts clients for a name upon connection.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

const timeout = 10 * time.Second

type client struct {
	Out  chan<- string // outgoing message channel
	Name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				select {
				case cli.Out <- msg:
				default:
					// Skip client if it's reading messages slowly.
				}
			}

		case cli := <-entering:
			clients[cli] = true
			cli.Out <- "Present:"
			for c := range clients {
				cli.Out <- c.Name
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.Out)
		}
	}
}

func handleConn(conn net.Conn) {
	out := make(chan string, 10) // outgoing client messages
	go clientWriter(conn, out)
	in := make(chan string) // incoming client messages
	go clientReader(conn, in)

	var who string
	nameTimer := time.NewTimer(timeout)
	out <- "Enter your name:"
	select {
	case name := <-in:
		who = name
	case <-nameTimer.C:
		conn.Close()
		return
	}
	cli := client{out, who}
	out <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli
	idle := time.NewTimer(timeout)

Loop:
	for {
		select {
		case msg := <-in:
			messages <- who + ": " + msg
			idle.Reset(timeout)
		case <-idle.C:
			conn.Close()
			break Loop
		}
	}

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func clientReader(conn net.Conn, ch chan<- string) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		ch <- input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
