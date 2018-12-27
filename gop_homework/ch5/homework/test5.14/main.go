//练习5.14： 使用breadthFirst遍历其他数据结构。
// 比如，topoSort例子中的课程依赖关系（有向图）,
// 个人计算机的文件层次结构（树），你所在城市的公交或地铁线路（无向图）。
package main

import "fmt"

func breadthFirst(f func(item string) []string, worklist *[]string) {
	seen := make(map[string]bool)
	for len(*worklist) > 0 {
		items := *worklist
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				fmt.Println(f(item))
				fmt.Println(*worklist)
				*worklist = append(*worklist, f(item)...)
			}
		}
	}
}

func main() {
	worklist := make([]string, 0)
	worklist = []string{
		"12345",
		"12343423323",
		"12324354trgfdwfwq",
		"3454rdfwetrewdf",
	}
	breadthFirst(func(item string) []string {
		return []string{"12"}
	}, &worklist)
	fmt.Println(worklist)
}
