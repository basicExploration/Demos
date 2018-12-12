package main

import (
	"testing"
)

func BenchmarkEathOther(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EathOther("12345", "54321")
	}
}

func TestEathOther(t *testing.T) {
	if k, v := EathOther("12345", "54321"); k {
		t.Log("ok", v)
	} else {
		t.Fatal("error", v)
	}
}



//go test
//PASS
//ok  	github.com/googege/gopl_homework/ch3/第三章练习题/练习3.12	0.008s

//go test -bench=. -cpu=4
//goos: darwin
//goarch: amd64
//pkg: github.com/googege/gopl_homework/ch3/第三章练习题/练习3.12
//BenchmarkEathOther-4   	30000000	        45.2 ns/op
//PASS
//ok  	github.com/googege/gopl_homework/ch3/第三章练习题/练习3.12	1.410s
