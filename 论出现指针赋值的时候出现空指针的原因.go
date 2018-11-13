// 根据这个可以延伸一下，之前经常出错或者是经常困惑的东西其实并不是引用类型例如slice map chan
// 其实一直以来经常出错或者是经常困惑的是指针。指针空的时候不能取*。跟引用类型没有一毛钱关系，引用类型的实体也是一个实体值。
package main

import "fmt"


type a struct {
	a2 string
	b1 *b
}
type b struct {
	c string
}

func TestValue() {
	var a1 a// 原因就是在于 a1.b1是个nil,指针的初始化就是nil,但是用nil是无法获取真实的内存地址的，所以就是出现提示nil pointer dereference.
	a1.b1 = &b{// 这个时候要做的就是给这个nil一个真实的内存地址即可。
		"",
	}
	a1.b1.c = "12"// 因为这个时候它已经被初始化了，不会是 nil 了。
	fmt.Println(a1.b1)
}

func main() {
	TestValue()
}
