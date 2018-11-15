//论go的new是谁的语法糖-论本来就是指针类型和取实体类型的指针
package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func main() {
	var a *tree       //a 原本就是指针类型
	var b tree        //b原本就是实体类型
	b1 := &b          //b1 原本就是实体类型的指针
	var c = new(tree) //new 就等于b1,它是先去实体再去指针的语法糖
	fmt.Printf("测试a c b1等不等于nil，%t,%t,%t,\n并且打印a,b,c,b1的值，%v,%v,%v,%v\n", a == nil, c == nil, b1 == nil, a, b, c, b1)
  a = b1
	fmt.Println("这个时候再看a的值：", a)
	fmt.Println("a的用途就是这样，是一种声明，不能直接赋值，因为它原本就是一个指针啊 ，一把钥匙而已，系统没有分给他房间，它当然不能打开门了")
}



//测试a c b1等不等于nil，true,false,false,
//并且打印a,b,c,b1的值，<nil>,{0 <nil> <nil>},&{0 <nil> <nil>},&{0 <nil> <nil>}
//这个时候再看a的值： &{0 <nil> <nil>}
//a的用途就是这样，是一种声明，不能直接赋值，因为它原本就是一个指针啊 ，一把钥匙而已，系统没有分给他房间，它当然不能打开门了
