package main

import (
	"fmt"
)

func countNumber(in, count int, out chan int) {
	for i := 0; i < count; i++ {
		out <- in
		in = in + count
	}
	close(out)
}

func Done(in chan int, done chan bool) {
	for v := range in {
		fmt.Println(v)
	}
	done <- true
}

func main() {
	a1 := make(chan int)
	b1 := make(chan bool)
	go countNumber(0, 20, a1)
	go Done(a1, b1)
	//time.Sleep(1e9)//设置阻塞也可以，只不过不知道设置多长时间，所以使用这种<-b1的方式，是让他自己阻塞然后听命令结束。
	<-b1 //设置它的目的就是告诉main你已经完成任务了不用再阻塞了。不然没有这个的话，main直接结束自己 上面两个线程都没有机会运行。
	fmt.Println("1")
}
