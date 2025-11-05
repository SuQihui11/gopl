package main

import (
	"fmt"
	"math"
)

func main() {

	v := math.MaxFloat64
	fmt.Println(v)

	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)    // "true"!

	var x float32 = 0.1
	fmt.Println(x)
	var y float32 = 0.2
	fmt.Println(x+y == float32(0.3))

	fmt.Printf("%.10f\n", x+y)

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"

}
