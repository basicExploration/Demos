package selectt

// select 的测试

import (
	"fmt"
)

func run(c, a chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-a:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	a := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		a <- 0
	}()
	run(c, a)
}
