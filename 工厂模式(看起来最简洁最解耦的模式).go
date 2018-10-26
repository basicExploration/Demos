package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("程序正式开始")
	dealEr()
	fmt.Println("你看我测试的最后结果成功的出来了吧")
}

func Sum() chan float64 {
	ch := make(chan float64)
	go func() {
		for i := 0.0; i < 10000000; i++ {
			ch <- i
		}
	}()
	return ch
}

func Suck(ch chan float64, done chan bool) {
	go func() {
		for {
			fmt.Println(<-ch)
		}
		done <- true
	}()

}

func sieve() {
	go func() {
		done := make(chan bool)
		ch := Sum()
		Suck(ch, done)
		<-done
	}()

}
func dealEr() {
	var t interface{}
	defer func() {
		if t = recover(); t != nil {
			fmt.Println("程序已经恢复", t)
		}
	}()
	sieve()
	time.Sleep(20 * 1e9)
	fmt.Println("panic 恢复正式开始", t)
}
