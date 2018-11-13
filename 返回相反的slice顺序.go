package main

import "fmt"

func main() {
	for k, v := range Reave([]int{1, 2, 3, 4, 5, 6}) {
		fmt.Println(k, v)
	}
}

// Reave 返回顺序相反的一个slice。
func Reave(sli []int) []int {
	for i, j := 0, len(sli)-1; i < len(sli)/2; i, j = i+1, j-1 {
		sli[i], sli[j] = sli[j], sli[i]
		fmt.Println(sli[i])
	}
	return sli
}
