package main

import "fmt"

func practice() (res int) {
	defer func() {
		if p := recover(); p != nil {
			res = 1111
		}
	}()
	//panic("this is a test")
	return
}

func main() {
	fmt.Printf("%d\n", practice())
}
