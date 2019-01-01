// 消费者生产者模式的经典例子
//主要是有两点

//1. 生产和和消费者这两个函数中（这个协程里）另外起一个匿名函数样式的协程然后一个生产者再没有提供参数的情况下返回一个产品
//然后由消费者直接使用
//2. 就是必须在两个函数中起一个 子函数协程

package main


import (
	"fmt"
	"time"
)

func main() {
	suck(sum())// 直接吃
	time.Sleep(1e9)
}

func sum() chan int {
	ch := make(chan int)
	go func() {// 要起一个协程其实在外面让sum起一个协程也可以不过这种方式看起来更加简洁。没有协程的话就不能使用通道就不能高并发异步了。
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	go func() {// 这里也是起一个协程，不然没办法接受通道传来的数据呀。
		for v := range ch {
			fmt.Println(v)
		}
	}()
}
