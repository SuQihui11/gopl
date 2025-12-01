package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type WordCounter int
type LineCounter int

func (w *WordCounter) Write(p []byte) (n int, err error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		*w++
	}
	return len(p), nil
}

func (l *LineCounter) Write(p []byte) (n int, err error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		*l++
	}
	return len(p), nil
}

type WriteCounter struct {
	w     io.Writer
	count int64 //因为计数器必须与包装的 Writer 共享 一个可变状态，从函数返回出去后，调用者要能看到计数值的实时变化。
	// 如果不用指针，调用者根本无法观察到计数的累加结果。
}

func (wc *WriteCounter) Write(p []byte) (n int, err error) {
	n, err = wc.w.Write(p)
	wc.count += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var n int64
	return &WriteCounter{w: w, count: n}, &n
}

func main() {
	//var wc WordCounter
	//wc.Write([]byte("hello world, 你好 世界 sqh"))
	//fmt.Println(wc) // 4
	//
	//var lc LineCounter
	//lc.Write([]byte("a\nb\nc\nsqh"))
	//fmt.Println(lc) // 3
	//
	//var w io.Writer
	//w = &wc
	//w.Write([]byte("dawdwadawdwada dwadadwadd adwadad"))
	//fmt.Printf("%[1]T, %[1]v\n", *w.(*WordCounter)) // 类型断言

	w, n := CountingWriter(os.Stdout)
	w.Write([]byte("Hello!\n"))
	fmt.Println("written:", *n)                      // 输出：written: 0
	fmt.Println("written:", w.(*WriteCounter).count) // 输出： 7
}
