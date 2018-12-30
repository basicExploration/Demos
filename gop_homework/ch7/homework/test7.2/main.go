//练习 7.2： 写一个带有如下函数签名的函数CountingWriter，传入一个io.Writer接口类型，
// 返回一个新的Writer类型把原来的Writer封装在里面和一个表示写入新的Writer字节数的int64类型指针
//
//func CountingWriter(w io.Writer) (io.Writer, *int64)
package main

import (
	"fmt"
	"io"
)

func ContingWriter(w io.Writer, p []byte) (io.Writer, *int64) {
	n, err := w.Write(p)
	if err != nil {
		fmt.Println(err)
	}
	a := int64(n)
	return w, &a
}