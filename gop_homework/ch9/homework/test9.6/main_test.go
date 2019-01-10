package main

import "testing"

func BenchmarkTT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TT()
	}
}
