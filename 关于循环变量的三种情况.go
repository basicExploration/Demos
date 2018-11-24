package main

import "fmt"

const T = 10

func main() {
	a := make([]func(), 0)
	// 第一种情况
	for i := 0; i < 10; i++ {
		a = append(a, func() { fmt.Println(i) })
	}
	for _, v := range a {
		v()
	}
	// 第二种情况

	for i := 0; i < 10; i++ {
		func() {
			fmt.Println(i)
		}()
	}

	// 第三种情况

	b := make([]func(), 0)
	// 第一种情况
	for i := 0; i < 10; i++ {
		t := i
		b = append(b, func() { fmt.Println(t) })
	}
	for _, v := range b {
		v()
	}

}
