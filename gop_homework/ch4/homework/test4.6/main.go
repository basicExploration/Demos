//练习 4.6： 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回
package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "2  49 d 2"
	st := []rune(s)
	result := make([]rune,0)
	for i := 0; i < len(st);  i++{

		if i +1 >= len(st) {
			result = append(result,st[i])
			break
		}
		if !unicode.IsSpace(st[i]){
			result = append(result,st[i])
		}
		if unicode.IsSpace(st[i]) && !unicode.IsSpace(st[i+1]) {
			result = append(result,st[i])
		}
	}
	fmt.Println(string(result))

}


