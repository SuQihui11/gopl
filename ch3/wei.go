package main

import "fmt"

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x) // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y) // "00000110", the set {1, 2}

	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}

	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}

	//medals := []string{"gold", "silver", "bronze"}
	//for i := uint8(2); i >= 0; i-- {
	//	fmt.Println(medals[i]) // "bronze", "silver", "gold"
	//}

	//var apples int32 = 1
	//var oranges int16 = 2
	//var compote int = apples + oranges // compile error

	//f := 3.141 // a float64
	//i := int(f)
	//fmt.Println(f, i) // "3.141 3"
	//f = 1.99
	//fmt.Println(int(f)) // "1"

	f := 1e100  // a float64
	i := int(f) // 结果依赖于具体实现
	fmt.Println(i)

}
