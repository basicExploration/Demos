//练习9.6: 测试一下计算密集型的并发程序(练习8.5那样的)会被GOMAXPROCS怎样影响到。
// 在你的电脑上最佳的值是多少？你的电脑CPU有多少个核心？
package main

import (
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	sy       sync.WaitGroup
	resulttt int64
)

func main() {
	TT()
}
func TT() {
	le := runtime.NumCPU()  // 分割的cpus数量
	dataLength := len(data) // data的数量
	t := dataLength / le
	sy.Add(le)
	for i := 0; i < le; i++ { // 假如是 32核的。每一个内核计算 dataLength / le
		var j []int64
		if i == le-1 {
			j = data[t*i:]
		} else {
			j = data[t*i : t*(i+1)]
		}
		go deal(j)
	}
	sy.Wait()
}

func deal(str []int64) {
	defer sy.Done()
	var result int64 = 0
	for _, v := range str {
		result += v
	}
	atomic.AddInt64(&resulttt, result)
}
