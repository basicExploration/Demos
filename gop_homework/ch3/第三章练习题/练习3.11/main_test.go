package main

import "testing"

var (
	s  float64 = 233333333333333333333333333333333333233333333333333333333333333333333333233333333333333333333333333333333333233333333333333333333333333333333333
	s2 string  = "233333333333333333333333333333333" +
		"333323333333333333333333333333333333333332333" +
		"3333333333333333333333333333333332333333333333333" +
		"333333333333333333333"
)

func BenchmarkCommabaBufferTestFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		commaBuffer(s)
	}
}

func BenchmarkCommabaBufferTestString(b *testing.B) { // 自动取对象的*
	for i := 0; i < b.N; i++ {
		commaBuffer(s2)
	}
}

//go test -bench=. -cpu=4
//goos: darwin
//goarch: amd64
//pkg: github.com/googege/gopl_homework/ch3/第三章练习题/练习3.11
//BenchmarkCommabaBufferTestFloat64-4   	  300000	      4728 ns/op
//BenchmarkCommabaBufferTestString-4    	 1000000	      1624 ns/op
//PASS
//ok  	github.com/googege/gopl_homework/ch3/第三章练习题/练习3.11	3.131s

//关于 方法 不论是定义在哪里的方法都是可以使用指针或者实际值取得这叫容错行

//关于指针方法使用对象，不需要麻烦的使用* 系统自动帮你取 * 然后取得对象，这叫特性，但是仅限于对象这个地方，所以这叫语法糖。此处的语法糖而已。

//什么叫对象 这里说的对象仅仅是指 struct，因为go里的对象就是struct而已。

//什么叫实现一个 接口 1  对象或者其它能实现方法的变量的实现 二  接口的嵌套。也叫实现，好好分析一下 接口没有实际用途那么如果使用一个 对象实现了
///那个新的对象 那么老的对象不也是实现了嘛，就是这个意思。
