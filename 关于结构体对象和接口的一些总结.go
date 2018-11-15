// 关于结构体对象 借口的一些总结。
package main

import "fmt"

type a struct {
	t string
}

func main() {
	// 对象的初始化一定要用new 引用的初始化一定要用make。空指针是nil只有在使出来对象的方法的时候才能使用但是会出错，不要用
	// 实现接口，是指针对象实现的还是值对象实现的不一样，是谁实现的就用谁不能混用。使用 var a1 *a a1是nil。
	//虽然可以使用 a1.TT 但是会出现bug不能使用 t对象，所以 var a1 *a 只能当做声明，而  var a1 a 就是可以真正的分配了内存了。
	//对于方法也是，是指针类型就用指针是值就用值，因为在没有接口的时候可以乱用但是在有接口的时候就不能了，所以要守规矩是谁的就用谁
	//就ok了。
	a1 := new(a)
	var dd ttt// 接口对象永远不能用指针类型。否则报错。
	dd = a1
	dd.TT()
}

func (a1 *a) TT() {
	fmt.Println(a1.t)
	fmt.Println("d")
}

type ttt interface {
	TT()
}
