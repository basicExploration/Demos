//练习 4.1： 编写一个函数，计算两个SHA256哈希码中不同bit的数目。（参考2.6.2节的PopCount函数。)
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	sha1 := sha256.Sum256([]byte("11!"))
	sha2 := sha256.Sum256([]byte("111"))
	fmt.Println(sha1, sha2)
	n := SHA256Count(sha1, sha2)
	fmt.Println("不同的书目是:", n)

}

func SHA256Count(shaOne, shaTwo [32]byte) int {
	var y int
	for i := 0; i < len(shaOne); i++ {
		if shaTwo[i] != shaOne[i] {
			y++
		}

	}
	return y
}

//[239 24 195 39 99 227 251 83 234 251 33 179 214 1 0 204 157 56 234 173 153 111 111 96 115 170 3 222 250 121 221 174] [246 224 161 226 172 65 148 90 154 167 255 138 138 170 12 235 193 42 59 204 152 26 146 154 213 207 129 10 9 14 17 174]
//不同的书目是: 31
// 一般情况下都是 全部不一样，好不容易有个31的。
