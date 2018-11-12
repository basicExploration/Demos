package main

import (
	"errors"
	"fmt"
)

func main() {
	if ok := IsRight("12345     ", "12344     "); ok != nil {
		fmt.Println("您输入的两个字符串并不是同一个字符串", ok)
	} else {
		fmt.Println("👌")
	}
}

//IsRight 判断s1和s2 是否是同一个字符串（字符相同但是顺序可以不一样）
func IsRight(s1, s2 string) error {
	var testValueY int
	if len(s1) != len(s2) {
		return errors.New("个数不同，不是同一个字符串")
	}
	byteS1, byteS2, le := []byte(s1), []byte(s2), len(s1)
	for i := 0; i < le; i++ {
		testValueX := 0
		for j := 0; j < le; j++ {
			if byteS1[i] != byteS2[j] {
				testValueX++
			}
		}
		if testValueX == le {
			return errors.New("存在不一样的字符码数")
		}
		testValueY++
	}
	if testValueY == le {
		return nil
	} else {
		return errors.New("存在不一样的字符码数")
	}
	return errors.New("未知错误")
}


