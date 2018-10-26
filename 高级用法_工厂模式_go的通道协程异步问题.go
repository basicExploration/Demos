package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := dd()
	go suck(ch1)
	time.Sleep(1e9)
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func dd() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}
