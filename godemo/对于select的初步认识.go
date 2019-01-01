  // 开始测试select

package main

import (
	"fmt"
	"time"
)

func main() {
	Suck()
	time.Sleep(1e9)
}

func Sum1() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {

			ch <- i
		}
	}()
	return ch
}

func Sum2() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {

			ch <- i
		}
	}()
	return ch
}

func Suck() {
	ch1 := Sum1()// 这个地方的线程其实就是main线程（因为一会要导入到main中。所以可以统称就是main线程中的一部分。）
	ch2 := Sum2()
	go func() {// 这才是开启了一个新的线程。
		for {
			select {
			case v := <-ch1:
				fmt.Println("这是ch1：", v)
			case v := <-ch2:
				fmt.Println("这是ch2", v)
			}
		}
	}()
}
