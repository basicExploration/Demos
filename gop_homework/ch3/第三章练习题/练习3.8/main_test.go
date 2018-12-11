package main

import "testing"

func Benchmarkcomp64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tt()
	}
}

func Benchmarkcom128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tt128()
	}
}
