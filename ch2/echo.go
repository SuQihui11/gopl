package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "Find the number")
var s = flag.String("s", " ", "space value")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *s))
	if !*n {
		fmt.Println("\n this is an enter")
	}
}
