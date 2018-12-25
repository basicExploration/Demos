package main

import "testing"

func BenchmarkSync(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Recover()
	}
}
