package main

import "fmt"

func main() {
	//var number int
	//fmt.Printf("%d\n", number)
	//
	//var name, sex string
	//fmt.Printf("name: %s, sex: %s\n", name, sex)
	//
	//var age, sqh = 12, "sqh"
	//fmt.Printf("age: %d, sqh: %s\n", age, sqh)

	//x := 1
	//p := &x         // p, of type *int, points to x
	//fmt.Println(*p) // "1"
	//*p = 2          // equivalent to x = 2
	//fmt.Println(x)  // "2"

	//var x, y int
	//fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"
	//var z *int
	//fmt.Println(z == nil)

	//var p = f()
	//var q = f()
	//fmt.Println(p == q)

	// 如果将指针作为参数调用函数，那将可以在函数中通过该指针来更新变量的值
	v := 1
	incr(&v)              // side effect: v is now 2
	fmt.Println(incr(&v)) // "3" (and v is 3)
	fmt.Println(v)

}

func f() *int {
	v := 1
	return &v
}
func incr(p *int) int {
	*p++ // 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
	return *p
}
