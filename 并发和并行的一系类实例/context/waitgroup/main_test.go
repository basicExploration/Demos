package main

import (
	"runtime"
	"testing"
)

func BenchmarkRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		read(doc1, runtime.NumCPU())
	}
}


func BenchmarkReadOneLine(b *testing.B){
	for i := 0;i < b.N ; i++ {
		readOneLine(doc2)
	}
}
