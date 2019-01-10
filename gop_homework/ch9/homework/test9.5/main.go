//练习 9.5: 写一个有两个goroutine的程序，两个goroutine会向两个无buffer channel反复地发送ping-pong消息。
// 这样的程序每秒可以支持多少次通信？
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var sy sync.WaitGroup
	sy.Add(100)
	ch := make(chan string)
	go func() { // 多发 一接受 就是说 可以让多个线程多发信息，让一个线程接受。
		for {
			ch <- "ping"
		}

	}()
	go func() {
		for {
			ch <- "pang"
		}

	}()
	go func() {
		defer sy.Done()
		for {
			<-ch
		}

	}()

	go func() {
		defer sy.Done() // 这个永远不会发生因为for不会停。
		for {
			start := time.Now()
			<- ch
			end := time.Now()
			fmt.Println(end.Sub(start))
		//	1.456µs
			//408ns
			//3.057µs
			//1.222µs
			//1.137265ms
			//12.005µs
			//1.321µs
			//34.037µs
			//1.45µs
			//397ns
			//428ns
			//3.613µs
			//1.163µs
			//1.554µs
			//410ns
			//421ns
		}

	}()

	sy.Wait()
}
