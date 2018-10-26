//本包主要是为了测试**闭包**这个特性

package closureExplain

import (
	"fmt"
)

func run() func(x int) int {// 注意 这个地方返回一个函数
	sum := 0
  return func(x int) int{//注意返回值，单个不需要括号，如果是多个是需要括号的。
	sum += x
	return sum
  }
}

func init() {
	a, b := run(), run()// run运行后返回一个func x
	
	a1 := a(10)// x 再次执行，返回一个sum 
	b1 := b(12)
	//这里我们就发现了，a b 接收到的两个run执行的返回值，也就是能绑定两个变量。
	fmt.Print(a1,b1)
}

