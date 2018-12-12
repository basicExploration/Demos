//练习 3.11： 完善comma函数，以支持浮点数处理和一个可选的正负号的处理。
package main

import (
	"bytes"
	"strconv"
)

func main() {
	t := "3232323223323232323232322332232323"
	var t1 float64 = 3443344343344343434343433443434343
	commaBuffer(t)
	commaBuffer(t1)

}

func commaBuffer(s interface{}) string {
	var buf bytes.Buffer
	b := 0
	if k, ok := s.(float64); ok {
		st := strconv.FormatFloat(k, 'f', 5, 64)
		if len([]byte(st)) <= 3 {
			return st
		}
		for _, v := range []byte(st) {
			buf.WriteByte(v)
			if b == 2 {
				buf.WriteByte(',')
				b = -1
			}
			b++
		}
	} else if k, ok := s.(string); ok {
		if len([]byte(k)) <= 3 {
			return k
		}
		for _, v := range []byte(k) {
			buf.WriteByte(v)
			if b == 2 {
				buf.WriteByte(',')
				b = -1
			}
			b++
		}
	}

	return buf.String()

}
