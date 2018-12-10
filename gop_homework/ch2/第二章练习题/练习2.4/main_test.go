package main

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(600)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountRewrite2(600)
	}
}
//pkg: github.com/googege/gopl_homework/ch2/第二章练习题/练习2.4
//BenchmarkPopCount-4    	2000000000	         0.37 ns/op
//BenchmarkPopCount2-4   	20000000	        69.7 ns/op
//PASS
//ok  	github.com/googege/gopl_homework/ch2/第二章练习题/练习2.4	2.267s
