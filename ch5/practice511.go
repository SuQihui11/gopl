package main

import (
	"fmt"
	"sort"
)

func main() {
	// 测试数据：添加calculus -> linear algebra的依赖，形成环
	var prereqsWithCycle = map[string][]string{
		"algorithms":     {"data structures"},
		"calculus":       {"linear algebra"}, // 原始依赖
		"linear algebra": {"calculus"},       // 新增：形成环！
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

	fmt.Println("测试环检测:")
	result, err := topoSortWithCycleDetection(prereqsWithCycle)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Println("拓扑排序结果:")
		for i, course := range result {
			fmt.Printf("%d:\t%s\n", i+1, course)
		}
	}
}

func topoSortWithCycleDetection(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	visiting := make(map[string]bool)

	var visitAll func(items []string) error
	visitAll = func(items []string) error {
		for _, item := range items {
			if !seen[item] {
				// 检测环形
				if visiting[item] {
					return fmt.Errorf("cycle detected: %s", item)
				}

				visiting[item] = true
				if err := visitAll(m[item]); err != nil {
					return err
				}
				visiting[item] = false
				order = append(order, item)
				seen[item] = true
			}
		}
		return nil
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	if err := visitAll(keys); err != nil {
		return nil, err
	}
	return order, nil
}
