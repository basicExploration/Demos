//说明一下，recover只能在defer中使用也就是defer然后来一个匿名函数，这样在里层栈中的panic在dd中发现了这个panic然后再defer中恢复即可。
//恢复了以后tt就可以继续进行了
package main

import "fmt"

func main() {
	dd(tt)
	fmt.Println("可以继续使用吗main")//恢复了以后这个逻辑就可以执行了。
}

func tt() {
	panic("测试")//只要有panic这个函数不论是否recover这个函数后面的东西都不会执行了。这个函数就废了。
	fmt.Println("可以继续使用吗tt")
}

func dd(f func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("恢复机制开始：", err)
		}
	}()
	f()
	fmt.Println("保护机制开始")
}

//##
//恢复机制开始： 测试
//可以继续使用吗main

// 事实证明：恢复是"可以的不过不能讲那个panic的函数恢复了，恢复的是忽略这个函数然后继续执行后面的逻辑"这个意思

//也就是说在哪一层恢复就哪一级别的逻辑可以继续执行，底层的逻辑这个recover管不了不然 那个 fmt.Println("可以继续使用吗tt")就可以执行了，事实上不能执行
//也就是验证了我说的。“哪一层恢复就哪一级别的逻辑可以继续执行，底层的逻辑这个recover管不了”
