package Parallelism



import "fmt"
//首先如果是并行处理的话，那么每一个go都会开启一个协程
//协程之间通信使用chanel 来通信。也就是简称的chan来通信。

func run (c chan int, d []int) {
	a := 0
	for _, value := range d {
		a += value
	}
	c <- a
}

func init() {
	a := []int{1, 2, 3 ,4 ,5 ,5}
	c := make(chan int)
	go run(c, a)
	go run(c, a)
	go run(c, a)
	go run(c, a)
	go run(c, a)
	go run(c, a)
	go run(c, a)
	go run(c, a)
	go run(c, a)
	go run(c, a)
	go run(c, a)
	go run(c, a)
	x, y, x1, x2, x3, x4, x5, x6, x7, x8, x9, x10 := <-c, <-c, <-c, <-c, <-c, <-c, <-c, <-c, <-c, <-c, <-c, <-c
	fmt.Print(x, y, x1, x2, x3, x4, x5, x6, x7, x8, x9, x10, )
}