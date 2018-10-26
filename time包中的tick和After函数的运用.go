package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	stop := time.After(1e9)// 它只能运行一次所以当它出现的时候可以设置为只要他出现就停止的这种选项
	tick := time.Tick(1e8)// 可以出现很多次而且是每隔设置的时间就出现一次。可以设置为报时器。
	for {
		select {
		case <-tick:
			fmt.Println(".")
		case <-stop:
			fmt.Println("boom")
			os.Exit(0)
		default:
			fmt.Println(".")
			time.Sleep(3e8)
		}
	}
}
