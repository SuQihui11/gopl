package main

var global *int

func main() {}

// 个x局部变量从函数f中逃逸了
func runAway() {
	var x int
	x = 12
	global = &x
}

func stay() {
	y := new(int)
	*y = 10
}
