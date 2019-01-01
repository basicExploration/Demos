package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	hash := md5.New() //返回的是一个已经实现过接口的了的struct，所以才可以直接使用。
	b := []byte{}
	n, err := hash.Write([]byte("666"))
	if n != 3 || err != nil {
		fmt.Println(err)
	}
	value := hash.Sum(b)
	hashString := fmt.Sprintf("%x\n", value)
	fmt.Println("输出的结果是:", hashString)

}
