//练习 9.2： 重写2.6.2节中的PopCount的例子，使用sync.Once，只在第一次需要用到的时候进行初始化。
// (虽然实际上，对PopCount这样很小且高度优化的函数进行同步可能代价没法接受)
package main

import (
	"fmt"
	"sync"
)

// pc[i] is the population count of i.
var pc [256]byte
var sy sync.Once

func init() {
	sy.Do(func() {// 只需要初始化一次
		for i := range pc {
			pc[i] = pc[i/2] + byte(i&1)
		}

	})

}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main(){
	fmt.Println(PopCount(12))
}