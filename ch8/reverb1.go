package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, text string, delay time.Duration) {
	fmt.Fprintln(c, "          "+strings.ToUpper(text))
	time.Sleep(delay)
	fmt.Fprintln(c, "          "+text)
	time.Sleep(delay)
	fmt.Fprintln(c, "          "+strings.ToLower(text))
}
func handleConnection(c net.Conn) {
	defer c.Close()
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		// 这里如果不加go，handleConn 函数被卡在 echo 里，无法回到 for input.Scan() 循环的开头
		go echo(c, scanner.Text(), 2*time.Second)
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
		go handleConnection(conn)
	}
}
