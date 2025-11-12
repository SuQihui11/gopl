package main

import "fmt"

// 5.10
func topoSortWithMap(m map[string][]string) map[int]string {
	var order = make(map[int]string)
	var index int
	seen := make(map[string]bool)

	var visitAll func(keys []string)
	visitAll = func(keys []string) {
		for _, key := range keys {
			if !seen[key] {
				seen[key] = true
				// 找出前置的链路
				visitAll(m[key])
				// 添加当前的结果是什么？
				order[index] = key
				index++
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

func main() {
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

	// 多次运行可能会得到不同的结果（遍历顺序不同）
	for i := 0; i < 3; i++ {
		fmt.Printf("\n运行第%d次:\n", i+1)
		result := topoSortWithMap(prereqs)
		for k, v := range result {
			fmt.Printf("%d:\t%s\n", k+1, v)
		}
	}
}
