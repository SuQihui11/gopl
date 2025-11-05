package main

import "fmt"

func main() {
	var runes []rune
	for _, r := range "Hello World!" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

	//var x, y []int
	//for i := 0; i < 10; i++ {
	//	y = appendInt(x, i)
	//	fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
	//	x = y
	//}

	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...) // append the slice x
	fmt.Println(x)      // "[1 2 3 4 5 6 1 2 3 4 5 6]"

	var y []string
	y = appendString(y, "Hello World!", "sqh", "love", "fq")
	fmt.Println(y)
}
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow.  Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling， for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}

func appendString(x []string, y ...string) []string {
	var z []string
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]string, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}
