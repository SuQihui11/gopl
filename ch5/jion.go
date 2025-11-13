package main

import (
	"fmt"
	"strings"
)

func Join(sep string, elems ...string) string {
	if len(elems) == 0 {
		return ""
	}
	if len(elems) == 1 {
		return elems[0]
	}

	// 计算总长度构建内存
	n := len(sep) * (len(elems) - 1)
	for i := range elems {
		n += len(elems[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(elems[0])
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}

	return b.String()
}

func main() {
	// 测试用例
	fmt.Printf("%q\n", Join(",", "a", "b", "c"))    // "a,b,c"
	fmt.Printf("%q\n", Join("-", "hello", "world")) // "hello-world"
	fmt.Printf("%q\n", Join(":", "only"))           // "only"
	fmt.Printf("%q\n", Join(", "))                  // ""
	fmt.Printf("%q\n", Join("|", "a", "", "c"))     // "a||c"
}
