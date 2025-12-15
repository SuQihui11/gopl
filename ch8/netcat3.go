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

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	conn.(*net.TCPConn).CloseWrite()
	<-done
}

//func main() {
//	ch := make(chan int, 1)
//	ch <- 1
//	close(ch) //正确，
//	end := <-ch
//	fmt.Println(end)
//}
