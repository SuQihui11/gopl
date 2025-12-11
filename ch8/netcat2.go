package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	// 客户端单独起一个goroutine等待，如果不用go，这里就会阻塞，因为一开始的时候服务器什么也不会输出（服务器的输出是依赖你的输入的，就造成了互相等）
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)

}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
