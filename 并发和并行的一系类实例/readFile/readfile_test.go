package main

import (
	"runtime"
	"testing"
)

func BenchmarkFind(b *testing.B) {
	for i := 0; i < b.N; i++ {

		find(datas)
	}
}

func BenchmarkF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findCon(runtime.NumCPU(), datas)
	}
}
