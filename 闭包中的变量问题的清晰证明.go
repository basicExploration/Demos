// 关于闭包中 闭包里的函数使用的是外部母函数中的变量的引用。
// 使用的是引用也是闭包的一种特性。
package main

import "fmt"

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

}

func lib() func() {
	i := 0
	return func() {
		i++
		fmt.Println(i, &i)

	}
}

// 1 0xc000012070
// 2 0xc000012070
// 3 0xc000012070
// 4 0xc000012070
// 分割线，因为b f不是一个函数实例所以这要证明他们的参数的引用值不是一个值
// 1 0xc000012078
// 2 0xc000012078
// 3 0xc000012078
// 4 0xc000012078

ps：关于另一个也是默认的是引用的一个变量是在循环体中的那个循环变量 参考：https://github.com/googege/goFiles/blob/master/关于匿名函数的几点.go
