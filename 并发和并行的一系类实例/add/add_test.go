package main

import (
	"fmt"
	"runtime"
	"testing"
)


func TestAdd(t *testing.T) {
	t.Log(add(num))
}

func TestNadd(t *testing.T) {
	t.Log(nadd(runtime.NumCPU(),num))
}


func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(num)
	}
	fmt.Println(len(num))// 经过计算 len是
}

func BenchmarkNadd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nadd(runtime.NumCPU(),num)
	}
	fmt.Println(len(num))
}


// 只有这种数量级的时候才能战胜单线程。 也就是 150k的时候 并发性能更好，如果是数量级非常小 比如 几千几万的时候 并发（并发什么时候都是不行的，
// 因为这是cpu bound）并行也是不行的。只有数量级非常大并且非常消耗计算量的时候 并行很厉害，但是也不会是跨越数量级的厉害。

//综上所述 cpu型 bound 并行 > 顺序执行 > 并发