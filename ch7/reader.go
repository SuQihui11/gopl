package main

import (
	"fmt"
	"io"
	"os"
)

type StringReader struct {
	s string
	i int
}

func NewReader(s string) *StringReader { return &StringReader{s, 0} }

func (r *StringReader) Read(p []byte) (n int, err error) {
	if r.i >= len(r.s) {
		return 0, io.EOF
	}
	n = copy(p, r.s[r.i:])
	r.i += n
	return
}

type LimitReader struct {
	r     io.Reader
	limit int64
}

func NewLimitReader(r io.Reader, n int64) *LimitReader { return &LimitReader{r, n} }

func (lr *LimitReader) Read(p []byte) (n int, err error) {
	if lr.limit <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > lr.limit {
		p = p[0:lr.limit]
	}
	n, err = lr.r.Read(p)
	lr.limit -= int64(n)
	return
}

func main() {

	// stringReader
	r := NewReader("<html><body>Hello</body></html>1111")

	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if n > 0 {
			os.Stdout.Write(buf[:n])
		}
		if err == io.EOF {
			fmt.Println("EOF\n")
			break
		}
	}

	// limitReader
	r1 := NewReader("<html><body>Hello</body></html>2222")
	lr := NewLimitReader(r1, 1024)

	buf = make([]byte, 1024)
	n, err := lr.Read(buf)
	fmt.Println(string(buf[:n]), err) // 输出：Hello <nil>

	n, err = lr.Read(buf)
	fmt.Println(string(buf[:n]), err) // 输出：  EOF
}
