//练习 4.9： 编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。
// 在第一次调用Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func wordfreq() {
	testMao := make(map[string]int)
	file, err := os.Open("./tt.txt")
	fmt.Println(err)
	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		testMao[input.Text()]++
		//fmt.Println(testMao)
	}
	for k,v := range testMao {
		fmt.Println(len(testMao))
		fmt.Printf("%s: %% %.3f",k,float32(v*100/len(testMao)))
	}

}

func main() {
	wordfreq()
}
