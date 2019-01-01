// 简单超时模式的demo
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	sieve()
}

func TimeOut(done chan bool) {
	go func() {
		time.Sleep(1e2)//此处设置的定时器，时间很短
		done <- true
	}()
}

func Creat(ch chan int) {
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i * 1000000000 * 232332//此处的计算稍微有点复杂，为了是能测试出来done的那个选项。
			//close(ch)
		}
	}()
}
func sieve() {
	done := make(chan bool)
	ch := make(chan int)
	TimeOut(done)
	Creat(ch)
	for {
		select {// 在协程通道模式中如果阻塞（所有协程都阻塞）并且程序没有结束就会panic，所以要么就停止工作要么就别阻塞（可以使用通道缓存）
		case <-done:
			fmt.Println("请求超时，立即断开。")
			os.Exit(0)
		case v := <-ch:
			fmt.Println(v)

		}
	}
}
