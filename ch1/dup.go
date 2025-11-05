package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	consts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		consts[line]++
	}
	for line, n := range consts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
