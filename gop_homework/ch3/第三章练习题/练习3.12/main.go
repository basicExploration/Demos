//练习 3.12： 编写一个函数，判断两个字符串是否是是相互打乱的，也就是说它们有着相同的字符，但是对应不同的顺序。
package main

import "fmt"

func main() {
	fmt.Println(EathOther("321", "123"))

}

func EathOther(s1, s2 string) (bool, string) {
	var ts int = 0
	var ty int = 0
	byteS1 := []byte(s1)
	byteS2 := []byte(s2)
	if len(byteS1) != len(byteS2) {
		return false, "len不相等"
	}
	for i := 0; i < len(byteS1); i++ {
		ts = 0
		for j := 0; j < len(byteS1); j++ {
			if byteS1[i] == byteS2[j] {
				ts++
			}

		}
		if ts > 0 {
			ty++
		}
	}
	if ty == len(byteS2) {
		return true, "nice"
	} else {
		return false, "有不相等的数据"
	}
	return false, ""

}
