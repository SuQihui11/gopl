package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
	}
}

func testTimeFormat() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))
	fmt.Println(now.Format("2006-01-02 03:04:05 PM"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000"))
	fmt.Println(now.Format(time.RFC3339))
}
