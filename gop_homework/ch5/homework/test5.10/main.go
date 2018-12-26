//练习5.10： 重写topoSort函数，用map代替切片并移除对key的排序代码。验证结果的正确性（结果不唯一）。
package main

import (
	"fmt"
	"sort"
)

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
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
	sort.Strings(keys)
	visitAll(keys)

	var t func()int
	t = func() int {
		return 12
	}
	t()
	return order
}


func main() {
	fmt.Println(topoSort(map[string][]string{
		"1": []string{"1", "2", "3"},
	}))
}
