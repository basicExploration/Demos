// 操作字符串的最佳方式：将它写入缓存中
package main

import (
	"bytes"
)

func main() {
	var buf bytes.Buffer// 一个Buffer struct实例。
	buf.WriteString("123")// 将字符串传到实例中
	buf.String()//读取缓存中的实例。
}
