//练习5.19： 使用panic和recover编写一个不包含return语句但能返回一个非零值的函数。
package main

import "fmt"

func tt(){
	defer func() {
		if e := recover();e != nil {
			fmt.Println(e)
		}
	}()
	dd()
}


func dd(){
	panic("12")
}

func main(){
	tt()
}