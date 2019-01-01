package main

import "testing"

func BenchmarkBubbleCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bubbleSort(num)
	}
}
func BenchmarkBubbleCountC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bubbleSort1(num)
	}
}

// 根据测试 main函数中的 end.Sub(start)的方法并不准确。

