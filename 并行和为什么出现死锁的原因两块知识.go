package main

import (
	"fmt"
	"time"
)

const T = 10

func main() {

	chan1 := make(chan int)
	chan2 := make(chan int)
	go func() {
		for i := 0; i < T; i++ {
			fmt.Println("发送者", time.Now().Nanosecond())
			chan1 <- i
		}
		close(chan1)
	}()
	for i := 0; i < T; i++ {
		go func(i int) {// 这里就是不断的让出现新的协程，这就是并行了。

			for c := range chan1 {
				fmt.Println("接受者和发送者", time.Now().Nanosecond())
				chan2 <- c
			}

		}(i)

	}
	for i := 0; i < 10; i++ {// 这里如果是跟传入的数据一样多 那么就不会出现 死锁，如果多了一个就会出现原因如下：
	// 本来应该是 没有了发送者然后接受者不需要一直阻塞应该是执行下面的东西或者是关闭了，但是这里的for会一直不停的接受
	//问题是没有了发送者了，程序就阻塞了，一直阻塞肯定不行啊 所以系统就崩溃了。如果刚好是10就没问题了。
		fmt.Println(<-chan2)
	}

	//们的和应该是经过通道的总个数，所以通道里的数据是一次性的取出来就没了不是值的复制

}
