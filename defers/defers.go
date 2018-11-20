package main

import "fmt"

func main() {
	fmt.Println(title(12))

}

func title(str int) (result int) {
	result = str
	defer func() { // defer 后面必须是函数。如果是这中返回值中已经确定了返回名称的形式，defer就可以更改return值，return的顺序是 
		//先执行return后面的东西 然后执行defer，然后执行return （返回，并停止。）
		result++ 
	}()
	return
}


func title2(str int) int {
	result := str
	defer func() { // defer 后面必须是函数。但是这里返回值并没有一个变量，这种形式，defer无法改变返回的结果。
		result++
	}()
	return result
}


//规则：后进先出
func run1() {
	defer fmt.Println("world")
	defer fmt.Println("wor")	
	      fmt.Print("hello")
}
// 规则： defer可以读取return后面的那个i值
func run2() (i int) {
	defer func() {
		i++ 
		fmt.Print(i)
	}()
	return 1
} 

// 规则：defer先定义后执行

func run3() {
	i := 0
	defer func(){fmt.Print(i)}()//答案是0, 因为这时候已经将i = 0 了所以是0
	i++
	defer func(){fmt.Print(i)}()//答案是1 ,i已经自增所以。

}
