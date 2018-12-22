package main

import "testing"

func BenchmarkT6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t6()
	}
}