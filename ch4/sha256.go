package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(c1, c2)
	start := time.Now()
	for i := 0; i < 100000; i++ {
		_ = comapre1(&c1, &c2)
	}
	fmt.Println(time.Since(start).Microseconds())
	start = time.Now()
	for i := 0; i < 100000; i++ {
		_ = compare2(&c1, &c2)
	}
	fmt.Println(time.Since(start).Milliseconds())
	//fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
	// Go语言对待数组的方式和其它很多编程语言不同，其它编程语言可能会隐式地将数组作为引用或指针对象传入被调用的函数。

	// 虽然通过指针来传递数组参数是高效的，而且也允许在函数内部修改数组的值；
	//但是数组依然是僵化的类型，因为数组的类型包含了僵化的长度信息。

}

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func comapre1(n1, n2 *[32]byte) int {
	count := 0
	for i, _ := range *n1 {
		count += popCount(n1[i], n2[i])
	}
	return count
}

func popCount(x, y byte) int {
	count := 0
	for i := 0; i < 8; i++ {
		xl := x % 2
		yl := y % 2
		if xl != yl {
			count++
		}
		x = x >> 1
		y = y >> 1
	}
	return count
}

func compare2(n1, n2 *[32]byte) int {
	count := 0
	for i, _ := range *n1 {
		diff := n1[i] ^ n2[i]
		count += popCount2(diff)
	}
	return count
}

func popCount2(x byte) int {
	count := 0
	for x > 0 {
		x &= x - 1
		count++
	}
	return count
}
