// go语言中的递归，这里是 递归效率最高的递归--- 尾递归
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	a := testWEI(100, 0)
	end := time.Now()
	fmt.Println("花费了多久：", end.Sub(start), a)
	// 这里就可以证明返回的最后的结果是最下层的那个函数层级的函数。
	// 尾递归跟递归的不同，因为它运行的栈不需要记录地址，而且也不会出现爆栈的现象，因为它一直在一条线上运行，结尾就是调用自己，而且只有
	//一个结尾，所以只调用一次，永远都是1.所以尾递归是单线结构。
}

func testWEI(n int, sum int) func(n int) {
	if n == 0 {
		fmt.Println(n, sum)
		return nil
	}
	sum += n
	fmt.Println(n)
	return testWEI(n-1, sum)
}
