package main

import (
	"errors"
	"fmt"
	"math"
)

func maxMin(nums ...int) (max, min int, err error) {
	if len(nums) == 0 {
		err = errors.New("no input")
		return 0, 0, err
	}

	min = math.MaxInt
	max = math.MinInt
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return max, min, nil
}

func main() {
	max, min, err := maxMin([]int{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10}...)
	fmt.Printf("max: %d, min: %d\n", max, min)
	fmt.Println(err)
}
