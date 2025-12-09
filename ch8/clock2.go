package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

var (
	port     = flag.String("port", ":8080", "The port to listen on")
	timeZone = flag.String("tz", "local", "The time zone to use")
)

func handleConnWithTimeZone(c net.Conn, location *time.Location) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().In(location).Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	flag.Parse()
	loc, err := time.LoadLocation(*timeZone)
	if err != nil {
		log.Fatal("Load time zone failed:", err)
	}

	listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConnWithTimeZone(conn, loc)
	}
}
