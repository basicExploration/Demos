// 关于闭包中 闭包里的函数使用的是外部母函数中的变量的引用。
// 使用的是引用也是闭包的一种特性。
package main

import (
	"fmt"
)

func main() {
	f := lib()
	b := lib()
	f()
	f()
	f()
	f()
	fmt.Println("分割线，因为b f不是一个函数实例所以这要证明他们的参数的引用值不是一个值")
	b()
	b()
	b()
	b()

	//------\

	var a1 a
	a1 = &byu{"12"}
	//a1 = byu{"12"} 这种方法为什么不行原因很简单就因为事先get的是指针类型。所以不匹配。不满足鸭子类型。
	//Cannot use 'byu{"12"}' (type byu) as type a in assignment Type does
	// not implement 'a' as 'get' method has a pointer receiver
	a1.get()
	
	
	//永远不要使用一个指针指向一个接口类型，因为它已经是一个引用类型了（这里就是指针）
	
	//错误演示
	// 这种方式就是错误的，在go里严禁取接口类型的指针。
	//
	//var a2 *a
	//*a2 = &byu{"12"}
	////a1 = byu{"12"} 
	//(*a2).get()
}

func lib() func() {
	i := 0
	return func() {
		i++
		fmt.Println(i, &i)

	}

}

type a interface {
	get()
}

type byu struct {
	value string
}

func (b1 *byu) get() {
	fmt.Println("n")
}

// 关于什么叫引用类型，1 指针 2 例如slice chan map中的那个带有cap len和 指针的 一种类似指针的类型。 所以可以理解为 引用对象大概就等于是 “指针”

ps：关于另一个也是默认的是引用的一个变量是在循环体中的那个循环变量 参考：https://github.com/googege/goFiles/blob/master/关于匿名函数的几点.go
