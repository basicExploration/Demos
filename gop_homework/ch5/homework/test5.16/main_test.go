package main

import (
	"testing"
)

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join([]string{"1","2","3","4"},"1","2","23","234343","2334","233323")
	}
}


//  为了实现同样的效果 strings.Join 必须这么写，那么就会变得很慢。但是自定义的Join并不会这样。

//
//go test -bench=. -cpu=4
//goos: darwin
//goarch: amd64
//pkg: github.com/googege/gopl_homework/ch5/homework/test5.16
//BenchmarkJoin-4     	10000000	       196 ns/op
//BenchmarkGoJoin-4   	 2000000	       707 ns/op
//PASS
//ok  	github.com/googege/gopl_homework/ch5/homework/test5.16	4.293s