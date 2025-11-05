package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetchToFile(url, ch, 1)
		go fetchToFile(url, ch, 2)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	}
	fmt.Printf("time elaseped: %.2f\n", time.Since(start).Seconds())
}

func fetchToFile(url string, channel chan<- string, number int) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		channel <- fmt.Sprintf("Exec_1.10: %v\n", err)
		return
	}

	fileName := strings.Split(url, ".")[1]
	if number == 1 {
		fileName += "_1.txt"
	} else {
		fileName += "_2.txt"
	}
	f, err := os.Create(fileName)
	if err != nil {
		channel <- fmt.Sprintf("Exec_2.10: %v\n", err)
		return
	}
	defer f.Close()

	n, err := io.Copy(f, resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		channel <- fmt.Sprintf("Exec_2.10: %v\n", err)
	}
	channel <- fmt.Sprintf("%.2f %7d %s [%d]", time.Since(start).Seconds(), n, url, number)
}
