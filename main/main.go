package main

import "fmt"

func run(a []int, b chan int) {
	total := 0
	for _, value := range a {
		total += value
	}
	b <- total
}

func big() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
func main() {
	big()
	b := make(chan int)
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	go run(a[:len(a)/2], b)
	go run(a[len(a)/2:], b)
	go run(a[:len(a)/2], b)
	go run(a[len(a)/2:], b)
	go run(a[:len(a)/2], b)
	go run(a[len(a)/2:], b)

	x, y, a1, b1, c, d := <-b, <-b, <-b, <-b, <-b, <-b
	fmt.Print(x, y, a1, b1, c, d)
}
