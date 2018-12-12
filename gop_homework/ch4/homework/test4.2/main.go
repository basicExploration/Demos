//练习 4.2： 编写一个程序，默认情况下打印标准输入的SHA256编码，并支持通过命令行flag定制，输出SHA384或SHA512哈希算法。
package main

import (
	"crypto/sha512"
	"fmt"
	"os"
)

func main(){
	for _,v := range os.Args[1:] {
		t := sha512.New().Sum([]byte(v))
		fmt.Println(t)
	}
}
