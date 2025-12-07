package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	fmt.Printf("type of w is %T\n", w) // type of w is <nil>， 此时接口值是nil，当然就没有值了
	w = os.Stdout                      // stdout 是一个变量，调用NewFile返回了一个file指针,而file实现了writer接口
	fmt.Printf("type of w is %T\n", w) // type of w is *os.File
	f := w.(*os.File)                  // 类型断言的结果是x的动态值，w的动态值此时是一个指向*os.File实例的地址
	fmt.Printf("type of f is %T\n", f)
	//c := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer
	//fmt.Printf("type of c is %T\n", c)
	main2()
	main3()
	fmt.Printf("========================================\n")
	test := (*os.File)(nil)
	demo(test)
}

type ByteCounter2 int

func (b *ByteCounter2) Write(p []byte) (n int, err error) {
	*b += ByteCounter2(len(p))
	return len(p), nil
}

func main2() {
	var w io.Writer
	w = os.Stdout
	rw := w.(io.ReadWriter) // success: *os.File has both Read and Write
	fmt.Printf("type of rw is %T\n", rw)
	//w = new(ByteCounter2)
	//rw = w.(io.ReadWriter) // panic: *ByteCounter has no Read method
}

func main3() {
	var w io.Writer = os.Stdout
	f, ok := w.(*os.File) // success:  ok, f == os.Stdout
	if !ok {
		fmt.Printf("type of f is %T\n", f)
	}
	b, ok := w.(*bytes.Buffer) // failure: !ok, b == nil
	if !ok {
		fmt.Printf("type of b is %T, val is %[1]v\n", b)
	}
}

func demo(w io.Writer) {
	// 【作用域 A】：外层的 w
	// 此时 w 的类型是接口类型 io.Writer
	fmt.Printf("外层 w 类型: %T\n", w)

	// if 语句开启了一个新的【作用域 B】
	// 这里的 := 声明了一个全新的变量，名字刚好也叫 w
	if w, ok := w.(*os.File); ok {
		// 【作用域 B 内部】：内层的 w
		// 此时 w 的类型是具体类型 *os.File
		// 在这里面使用 w，都指代这个新的 *os.File 变量
		fmt.Printf("内层 w 类型: %T (原有 w 被遮蔽)\n", w)
		w.Close() // 调用的是 *os.File 的 Close
	}

	// 【回到作用域 A】
	// 出了 if 块，内层的 w 销毁了。
	// 这里的 w 依然是原来的接口类型 io.Writer，没有被修改
	fmt.Printf("回到外层 w 类型: %T\n", w)
}
