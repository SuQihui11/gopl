package main

import "fmt"

func main() {
	var i int
	var x int32
	i = 12
	x = int32(i)
	fmt.Println(x)
	y := -5 % -2
	z := -5 % 2
	fmt.Println(y, z)

	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // "255 0 1"

	var j int8 = 127
	fmt.Println(j, j+1, j*j) // "127 -128 1"

	s := 1
	q := -1
	z = s &^ q
	fmt.Println(z)
}
