//练习 1.1： 修改echo程序，使其能够打印os.Args[0]，即被执行命令本身的名字。
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ { // Args[0]就是执行的命令本身。
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("打印结果：", s)
}
// go run main.go 56 12 14
///var/folders/x7/dfv542p515dd6h37f69d506c0000gn/T/go-build524123334/b001/exe/main 56 12 14

//可以看出 go run 的命令，会生成一个临时的文件。
