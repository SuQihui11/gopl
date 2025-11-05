package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	//s, sep := "", ""
	//for _, arg := range os.Args[1:] {
	//	// 每次循环迭代字符串 s 的内容都会更新。
	//	//+= 连接原字符串、空格和下个参数，产生新字符串，并把它赋值给 s。
	//	//s 原来的内容已经不再使用，将在适当时机对它进行垃圾回收。
	//	s += sep + arg
	//	sep = " "
	//}
	str := "1dwda eafafw dfwafawfd dfawfafwwafa"
	n := 100000
	printArgs1(n, str)
	printArgs2(n, str)
}

func TimeConsuming(tag string) func() {
	start := time.Now().UnixNano()
	return func() {
		end := time.Now().UnixNano()
		fmt.Printf("%s cost time:%d\n", tag, end-start)
	}
}

func printArgs1(n int, str string) {
	defer TimeConsuming("printArgs1")()
	res, seq := "", ""
	for i := 0; i < n; i++ {
		res += seq + str
		seq = " "
	}
	_ = res
}

func printArgs2(n int, str string) {
	defer TimeConsuming("printArgs2")()
	args := make([]string, n)
	for i := 0; i < n; i++ {
		args[i] = str
	}
	_ = strings.Join(args, " ")
}
