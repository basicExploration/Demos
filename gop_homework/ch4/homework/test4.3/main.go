//练习 4.3： 重写reverse函数，使用数组指针代替slice。
package main

import "fmt"

func reverse1(s [5]int) { // 这么写的时候  s 已经是复制值了，所以失败。
	st := &s
	for i, j := 0, len(*st)-1; i < j; i, j = i+1, j-1 {
		(*st)[i], (*st)[j] = (*st)[j], (*st)[i]
		//fmt.Println(*st)
	}
}

func reverse2(s *[5]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]

	}
}
func main() {
	t := [5]int{1, 443, 223, 2, 0}
	reverse1(t)
	fmt.Println(t)
	reverse2(&t)
	fmt.Println(t)
}
