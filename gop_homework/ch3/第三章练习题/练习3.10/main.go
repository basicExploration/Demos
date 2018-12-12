//练习 3.10： 编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作。
package main

import (
	"bytes"
	"fmt"
)

func main() {
	t := []string{"111", "2222277777"}
	for i := 0; i < len(t); i++ {
		fmt.Printf("  %s\n", comma(t[i]))
		fmt.Println("2:::", commaBuffer([]byte(t[i])))
	}

}

//!+ 速度慢是因为递归。
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]

}

func commaBuffer(s []byte) string {
	var buf bytes.Buffer
	b := 0
	if len(s) <= 3 {
		fmt.Println(string(s))
		return string(s)
	}
	for _, v := range s {
		buf.WriteByte(v)
		if b == 2 {
			buf.WriteByte(',')
			b = -1
		}
		b++
	}
	return buf.String()

}
