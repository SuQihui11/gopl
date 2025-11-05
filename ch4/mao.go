package main

import (
	"fmt"
	"sort"
)

func main() {
	//ages := make(map[string]int) // mapping from strings to ints
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	//delete(ages, "alice") // remove element ages["alice"]
	fmt.Println(ages)
	//_ = &ages["bob"] // compile error: cannot take address of map element
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	emap := make(map[string]int)
	emap["alice"] = 31
	var emap2 map[string]int
	emap2["charlie"] = 34
}
