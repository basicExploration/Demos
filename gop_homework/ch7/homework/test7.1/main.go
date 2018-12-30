//练习 7.1： 使用来自ByteCounter的思路，实现一个针对对单词和行数的计数器。你会发现bufio.ScanWords非常的有用。
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(T())
}
func T() int {
	var d int
	var t = func() {
		d++
	}
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords) // bufio.ScanWords 等这些函数符合SplitFunc函数的特征，那么就是符合这个类型。

	for {
		if scan.Scan() {
			fmt.Println(scan.Text())
			t()

		}
		if scan.Err() == io.EOF {
			break
		}
		if scan.Text() == "Q" {
			break
		}
	}
	//
	return d
}
