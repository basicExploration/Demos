// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//工厂模式就是每次都是返回一个已经加工好的值。
// 工厂模式可以很好的解耦，go的协程的工厂模式就是比如三个函数每个函数中来一个 go func(){} 起一个协程即可，然后
// 因为没有关闭通道然后接受者可能造成了阻塞，（接受者的阻塞来自于发送者的没有准备好，当最后一个了没有close就会发生没有准备好的情况就会阻塞导致死机）
//这个时候 就会panic(gc也会关闭通道但是那是不影响程序运行的情况下)所以需要我们用recover来恢复panic的进程（recover记得要在defer中才能恢复。）
// 用defer func(){}即可。
package main

import (
	"fmt"
)

// Send the sequence 2, 3, 4, ... to returned channel
func generate() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch // 返回的是数字
}

// Filter out input values divisible by 'prime', send rest to returned channel
func filter(in chan int, prime int) chan int {// 返回的是筛选的结果。
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func sieve() chan int {// 将数字发生器和筛选器结合在一起。
	out := make(chan int)
	go func() {
		ch := generate()
		for {
			prime := <-ch
			ch = filter(ch, prime)
			out <- prime
		}
	}()
	return out
}

func main() {
	primes := sieve()
	for {
		fmt.Println(<-primes)
	}
}
