package main

import "fmt"

type a int

func (a1 a) get() {

}

type b interface {
	get()
}

func main() {
	var a1 a
	var b1 b
	b1 = a1
	fmt.Println(b1)
	b1.get()
}

// 证明了 所有类型都可以实现接口，甚至是 接口类型。（只需要这个接口和那个接口定义的方法一直，就可以直接赋值，不过
// 没什么意义）。但 指针类型不可以，因为 指针类型不能作为 type后面的底层，所以它就无法实现接口了。
