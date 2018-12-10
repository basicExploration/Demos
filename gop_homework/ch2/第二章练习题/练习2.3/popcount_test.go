package main

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(1000)
	}
}

func BenchmarkPopCountCycle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountCycle(1000)
	}
}

//go test -cpu=4 -bench=.
//goos: darwin
//goarch: amd64
//pkg: github.com/googege/gopl_homework/ch2/第二章练习题/练习2.3
//BenchmarkPopCount-4        	2000000000	         0.37 ns/op
//BenchmarkPopCountCycle-4   	100000000	        18.8 ns/op
//PASS
//ok  	github.com/googege/gopl_homework/ch2/第二章练习题/练习2.3	2.698s
