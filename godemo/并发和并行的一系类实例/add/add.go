package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func add(n []int) int64 {
	var result int64

	for _, v := range n {
		result += int64(v)
	}
	return result
}

func nadd(goruntimes int, n []int) int64 {
	var tt int64
	totalNumber := len(n)              // 计算一共的量
	perNum := totalNumber / goruntimes // 计算每一个应该分多少
	var sy sync.WaitGroup
	sy.Add(goruntimes)
	for i := 0; i < goruntimes; i++ {
		go func(i int) {
			start := perNum * i
			end := perNum + start
			if i == goruntimes-1 {
				end = totalNumber
			}
			var t int64
			for _, v := range n[start:end] {
				t += int64(v)
			}
			atomic.AddInt64(&tt, int64(t))
			sy.Done()
		}(i)
	}
	sy.Wait()
	return tt
}

func main(){
	fmt.Println(add(num))
	fmt.Println(nadd(runtime.NumCPU(),num))
}
