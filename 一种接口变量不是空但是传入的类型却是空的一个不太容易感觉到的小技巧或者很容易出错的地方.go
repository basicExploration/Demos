package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = false

func main() {
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	fmt.Println(buf == nil)
	f(buf) // NOTE: subtly incorrect!
	fmt.Println(buf.String())
	if debug {
		// ...use buf...
	}
}

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
	// ...do something...
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
// 当 debug 是false的时候，那么 buf就是一个 以buffer类型的一个nil值，但是它赋予了io.Writer,并且buffer类型实现了io.Writer
// 所以是可以实现继承的，但是它的值却是一个nil，但是与此同时 这个out变量的值可不是nil 它是一个 有着类型但是指向是nil的一个非nil使用的时候肯定是会
//报bug的。
