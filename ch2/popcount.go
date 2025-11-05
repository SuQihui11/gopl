package main

import (
	"fmt"
	"time"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
	var count int
	for i := uint64(0); i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}

func PopCount3(x uint64) int {
	var count int
	for x != 0 {
		if x&1 != 0 {
			count++
		}
		x >>= 1
	}
	return count
}

func PopCount4(x uint64) int {
	var count int
	for x != 0 {
		x = x & (x - 1)
		x -= 1
		count++
	}
	return count
}

func main() {
	arr := make([]int, 200000000)
	for i := range arr {
		arr[i] = i + 1
	}
	start := time.Now()
	for i := range arr {
		PopCount(uint64(arr[i]))
	}
	fmt.Printf("%ds elapsed\n", time.Since(start).Milliseconds())

	start = time.Now()
	for i := range arr {
		PopCount2(uint64(arr[i]))
	}
	fmt.Printf("%ds elapsed\n", time.Since(start).Milliseconds())

	start = time.Now()
	for i := range arr {
		PopCount3(uint64(arr[i]))
	}
	fmt.Printf("%ds elapsed\n", time.Since(start).Milliseconds())

	start = time.Now()
	for i := range arr {
		PopCount4(uint64(arr[i]))
	}
	fmt.Printf("%ds elapsed\n", time.Since(start).Milliseconds())
}
