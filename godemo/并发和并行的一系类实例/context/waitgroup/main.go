package main

import (
	"sync"
	"time"
)

var doc1 = []int{1, 2, 3, 4, 5, 3, 6, 5, 4, 4, 4, 33, 3, 3, 3, 3, 3, 3}
var doc2 = []int{1, 2, 3, 4, 5, 3, 6, 5, 4, 4, 4, 33, 3, 3, 3, 3, 3, 3}

func read(doc []int, cpus int) []int {
	result := make([]int, 0)
	ch := make(chan int, len(doc))
	for i := 0; i < len(doc); i++ {
		ch <- doc[i]
	}
	close(ch)
	var sy sync.WaitGroup
	sy.Add(cpus)
	for i := 0; i < cpus; i++ {
		go func() {
			for v := range ch {
				read1(&v)
				result = append(result, v)
			}
			defer sy.Done()
		}()
	}
	sy.Wait()
	return result
}

func read1(s *int) {
	time.Sleep(1e6)
	*s += 100
}

func readOneLine(doc []int) []int {
	result := make([]int,0)
	for _,v := range doc {
		read1(&v)
		result = append(result,v)
	}
	return result
}

//go test -bench=. -cpu=1
//goos: darwin
//goarch: amd64
//pkg: github.com/googege/goFiles/并发和并行的一系类实例/context/waitgroup
//BenchmarkRead        	     200	   7139032 ns/op
//BenchmarkReadOneLine 	      50	  25281216 ns/op
//PASS
//ok  	github.com/googege/goFiles/并发和并行的一系类实例/context/waitgroup	3.446s
//ThomasHukes-MacBook-Air:waitgroup ThomasHuke$ go test -bench=. -cpu=4
//goos: darwin
//goarch: amd64
//pkg: github.com/googege/goFiles/并发和并行的一系类实例/context/waitgroup
//BenchmarkRead-4          	     200	   7225286 ns/op
//BenchmarkReadOneLine-4   	      50	  25757595 ns/op
//PASS
//ok  	github.com/googege/goFiles/并发和并行的一系类实例/context/waitgroup	3.506s


//遇到阻塞比较严重的可以看得出来 这种io bound的 情况 异步是好不少的。但是只能说是量大的时候 量少的时候 非异步更快！！！！（并且并行和并发并没有什么大的区别。。。）
