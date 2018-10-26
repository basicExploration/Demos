package main

import (
	"fmt"
	"os"
)

func tel(ch chan int, done chan bool) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	done <- true// 当ch已经传递完成了就开启done。
}

func main() {
	var ok = true
	ch := make(chan int)
	done := make(chan bool)
	go tel(ch, done)
	for ok {
		select {// 接受两个线程，当接收到done的时候直接退出。
		case v := <-ch:
			fmt.Println(v)
		case <-done:
			os.Exit(0)
		}
	}
}
