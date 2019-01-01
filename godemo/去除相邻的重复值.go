// 去除相邻的重复值
package main

import "fmt"

func main() {
	fmt.Println(tt("11111111111111fewfefd44334434343434"))
}
func tt(str string) string {
	result := make([]byte, 0)
	byteStr := []byte(str)
	for i := 0; i < len(byteStr)-1; i++ {
		if byteStr[i+1] == byteStr[i] {// 只要前后一样那么就continue
			continue
		} else {
			result = append(result, byteStr[i])// 一旦不一样了就把这个数字append给result slice。
		}
	}
	return string(append(result,byteStr[len(byteStr)-1:]...))
  // 因为最后一个数字是无法处理的，并且最后一个数字无论如果都是要加上的，所以直接
  //添加上即可。
}
