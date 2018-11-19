//关于匿名函数的几点意见，有匿名函数的地方容易有闭包
// 关于闭包闭包里使用的母函数中的局部变量会一直存活在闭包中
//并且这个变量还是一直是引用变量，就是是用的地址，这就是闭包的特性之一。
package main

import (
	"os"
)

func main() {
	lib()
}

func lib() {
	var rmdirs []func()
	dirs := tempDirs()
	for i := 0; i < len(dirs); i++ {
		j := i
		// 这里因为并不牵涉到闭包所以就是正常的每次循环的i
		// 这里之所以是使用一个局部变量是因为想记录每次的i值因为闭包中的i一直是引用值
		//但是如果使用j这个局部变量的话，那么每次使用的就是实际值了。
		os.MkdirAll(dirs[i], 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dirs[j]) // NOTE: incorrect!
		})
	}
	for _, v := range rmdirs {
		v()
	}
}

func tempDirs() []string {
	return []string{"1", "2", "3", "4", "5"}
}
