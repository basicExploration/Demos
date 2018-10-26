package slice

import "fmt"

func run() {
	a := []int{1, 2, 3, 4}
	a1 := a[1:3]//表示从2到3的数字
	a2 := append(a, 1, 2, 3, 3)//这里说的就是说我们将 1233 传入给a的复制品然后返回一个新的切片。
	fmt.Print(a2, a1)
	
}