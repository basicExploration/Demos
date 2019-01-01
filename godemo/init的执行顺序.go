package main

import "fmt"

var t int

func main() {
	d()
	fmt.Println("谁先开始main")
}
func init() { // 在这个函数里预先经很多函数先预执行很不错的选择。
	t = 12
	fmt.Println("谁先开始:init")
}

func d() {
	fmt.Println(t)
}

### 

谁先开始:init
12
谁先开始main
