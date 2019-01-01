package main

// 主要目的是为了测试并发的channel模块

import (
	"fmt"
)

func sum(a []int, c chan int) {
	var sum int
	for _, value := range a {
		sum += value
	}
	c <- sum

}

func main() {
	a := []int{
		1,
		2,
		3,
		4,
		5,
	}

	c := make(chan int, 112)

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y)
}
