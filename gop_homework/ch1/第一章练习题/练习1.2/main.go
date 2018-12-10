//练习 1.2： 修改echo程序，使其打印每个参数的索引和值，每个一行。
package main

import (
	"fmt"
	"os"
)

func main() {

	for i := 1; i < len(os.Args); i++ {
		fmt.Println(i, ":", os.Args[i])
	}

}

//go run main.go 12 23 123
//1  :  12
//2  :  23
//3  :  123
