package main

import (
	"fmt"
)

func main() {
	//f := squares()
	//
	////
	//fmt.Println(f()) // "1"
	//fmt.Println(f()) // "4"
	//fmt.Println(f()) // "9"
	//fmt.Println(f()) // "16"
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(item []string)
	visitAll = func(item []string) {
		for _, item := range item {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	visitAll(keys)
	return order
}

// prereqs记录了每个课程的前置课程
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}
