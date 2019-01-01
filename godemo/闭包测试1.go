// 来测试闭包: 闭包使用的外部变量是引用值 不论是值还是指针一律使用引用，所以i才会多调用几次一直增加没有脱离那个栈。
package main

import "fmt"

func main() {
	i := 0
	a := testBibao(i)
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	j := 1
	b := testBibao(j)
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}

func testBibao(i int) func() int {
	return func() int {
		i++
		return i
	}
}


1
2
3
4
5
2
3
4
5
6



package main

import (
	"fmt"
)

func main() {
	errHnadler(12, func(i int) {
		a, b := 10, 0
		fmt.Println(*(&a) / *(&b))
	})
	fmt.Println("如果恢复成功的话，这条信息是可以打印出来的。")
}

type tt func(int)

func errHnadler(in int, fn tt) {
	defer func() {
		if err := recover(); err != 0 {
			fmt.Println(err)
		}
	}()
	fn(in)
	fmt.Println("ceshi")
}
