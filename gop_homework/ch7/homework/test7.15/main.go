//练习 7.15： 编写一个从标准输入中读取一个单一表达式的程序，用户及时地提供对于任意变量的值，
// 然后在结果环境变量中计算表达式的值。优雅的处理所有遇到的错误。
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main(){
	data,err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	//这个练习是什么鬼意思？
}

