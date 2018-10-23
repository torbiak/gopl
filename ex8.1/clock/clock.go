// Clock creates a concurrent clock server that can accept a port number
// and can be configured to a time zone through checking the environment
// variable TZ.  TZ should be set to a region used in the IANA time database
// accessible for reference here: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
//
// Three servers that work for the demo are:
// TZ=America/New_York ./clock_server -port 8000 &
// TZ=America/Chicago ./clock_server -port 8010 &
// TZ=America/Los_Angeles ./clock_server -port 8020 &
//
// Build Note: the folder is named clock and so is the program, therefore
// when building use `go build -o clock_server` to create a unique executable.

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var port = flag.Int("port", 8000, "localhost port for serving the clock")

func main() {
	flag.Parse()
	fmt.Println("Port is:", *port)
	addr := net.JoinHostPort("localhost", fmt.Sprint(*port))

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	// Set the timezone.
	z := setZone()

	log.Println("Clock server listening on:", addr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g. Connection aborted
			continue
		}
		go handleConn(conn, z) //handle one connection at a time
	}
}

func setZone() *time.Location {
	//init the loc
	zone := os.Getenv("TZ")
	if zone == "" {
		log.Println("TZ not set, using default: America/Chicago")
		zone = "America/Chicago"
	}
	loc, err := time.LoadLocation(zone)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Loaded zone:", zone)

	return loc
}

func handleConn(c net.Conn, z *time.Location) {
	defer c.Close()

	for {
		_, err := io.WriteString(c, time.Now().In(z).Format("15:04:05\n"))
		if err != nil {
			return //e.g. Client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
