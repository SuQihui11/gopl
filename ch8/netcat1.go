package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	// 2. 将连接中的数据（conn），直接拷贝到标准输出（os.Stdout）
	// 这样你就能在终端看到服务器发来的时间了
	if _, err = io.Copy(os.Stdout, conn); err != nil {
		log.Fatal(err)
	}
}
