package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, text string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(text))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", text)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(text))
}
func handleConnection(c net.Conn) {
	defer c.Close()
	// 1.定义 WaitGroup 用来计数活跃的 echo goroutine
	var wg sync.WaitGroup

	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		// 2. 每次启动 echo 前，计数加 1
		wg.Add(1)
		go func(input string) {
			defer wg.Done()
			echo(c, input, 1*time.Second)
		}(scanner.Text())
		// 这里如果不加go，handleConn 函数被卡在 echo 里，无法回到 for input.Scan() 循环的开头
		//go echo(c, scanner.Text(), 2*time.Second)
	}

	// 4. 循环结束意味着客户端关闭了写入 (EOF)。
	// 此时我们不能立即关闭连接，必须等待所有 echo 喊完。
	wg.Wait()

	if tcpConn, ok := c.(*net.TCPConn); ok {
		tcpConn.CloseWrite()
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
