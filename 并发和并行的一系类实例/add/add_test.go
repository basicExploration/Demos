package main

import (
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
}

func BenchmarkNadd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nadd(runtime.NumCPU(),num)
	}
}

