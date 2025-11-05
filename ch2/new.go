package main

import "fmt"

func main() {
	p := new(int)
	q := new(int)
	fmt.Println(*p)
	fmt.Println(*q)
	fmt.Println(p == q)
}
