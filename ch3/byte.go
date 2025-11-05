package main

import (
	"fmt"
	"math"
)

const (
	KB = 1 << (10 * iota)
	MB = KB << 10
	GB = MB << 10
	TB = GB << 10
	PB = TB << 10
	EB = PB << 10
	ZB = EB << 10
	YB = ZB << 10 // 发生了溢出
)

func main() {
	fmt.Printf("%T %d %d\n", KB, MB, GB)
	fmt.Println(math.MaxInt)
	fmt.Println(ZB < math.MaxInt)

	// YiB/ZiB是在编译期计算出来的，并且结果常量是1024，是Go语言int变量能有效表示的
	fmt.Println(YB / ZB)
	//var x float32 = math.Pi
	//var y float64 = math.Pi
	//var z complex128 = math.Pi
	//
	//const Pi64 float64 = math.Pi
	//
	//var x float32 = float32(Pi64)
	//var y float64 = Pi64
	//var z complex128 = complex128(Pi64)

	var f float64 = 212
	fmt.Printf("%T, %[1]g\n", (f-32)*5/9)     // "100"; (f - 32) * 5 is a float64
	fmt.Printf("%T, %[1]v\n", 5/9*(f-32))     // "0"; 5/9 is an untyped integer, 0
	fmt.Printf("%T, %[1]g\n", 5.0/9.0*(f-32)) // "100"; 5.0/9.0 is an untyped float

	fmt.Printf("%T\n", 0)      // "int"
	fmt.Printf("%T\n", 0.0)    // "float64"
	fmt.Printf("%T\n", 0i)     // "complex128"
	fmt.Printf("%T\n", '\000') // "int32" (rune)

}
