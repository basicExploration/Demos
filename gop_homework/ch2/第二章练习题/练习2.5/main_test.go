package main





import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(600)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountRewrite3(600)
	}
}


//go test -bench=. -cpu=4
//goos: darwin
//goarch: amd64
//pkg: github.com/googege/gopl_homework/ch2/第二章练习题/练习2.5
//BenchmarkPopCount-4    	2000000000	         0.38 ns/op
//BenchmarkPopCount2-4   	200000000	         7.86 ns/op
//PASS
//ok  	github.com/googege/gopl_homework/ch2/第二章练习题/练习2.5	3.168s