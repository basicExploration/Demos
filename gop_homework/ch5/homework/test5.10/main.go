//练习5.10： 重写topoSort函数，用map代替切片并移除对key的排序代码。验证结果的正确性（结果不唯一）。
// 深度优先算法
// 广度 优先算法
//todo：未能解决问题。
package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"log"
)

//深度
func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func([]string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] { // 如果还没有定义过得 bool初始值是false
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for _, v := range m {
		keys = append(keys, v...)
	}
	visitAll(keys)

	//var t func() int
	//t = func() int {
	//	return 12
	//}
	//t()
	return order
}

//--------
//广度

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-breadthFirst

//!+crawl
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	t := map[string][]string{
		"1": {"1", "2", "3", "332"},
		"2": {"2", "23", "32", "554"},
	}
	fmt.Println(topoSort(t))
}

//func tt(){
//
//	t := []int{
//		1,3,3,5,3,
//	}
//	d := [][]int{
//		{1,3,23,},
//	}
//	for _,v := range t  {
//		for _,d := range d[v]{
//			fmt.Println(d)
//		}
//	}
//}
