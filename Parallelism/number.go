package Parallelism

import "fmt"

func init() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Print(<-c)
	fmt.Print(<-c)
}
