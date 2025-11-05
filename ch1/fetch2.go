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
	defer timeConsuming("fetch url")()
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		s, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("\n\n written_size: %d\n\n", s)
		fmt.Printf("\n\n resp_code: %s\n", resp.Status)
	}
}
func timeConsuming(tag string) func() {
	start := time.Now().UnixNano()
	return func() {
		end := time.Now().UnixNano()
		fmt.Printf("%s cost time:%d\n", tag, end-start)
	}
}
