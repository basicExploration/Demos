package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var i int // 测试的次数
	value := "💪"
	var byteValue  = []byte(value)
	file, _ := os.Open("./go.mod")
	file2, err := os.OpenFile("./dd.txt", os.O_WRONLY|os.O_CREATE, 0666) // 这里一定要写成这个是以为 os.O_WRONLY|os.O_CREATE 没有就建一个这个关键点
	if err != nil {
		fmt.Println("在创建要写入的文件的时候出错", err)
	}
	defer file.Close()
	defer file2.Close()
	reader := bufio.NewReader(file) // 读文件
	writeer := bufio.NewWriter(file2)
	for {
		byte, err := reader.ReadByte() // 将reader具体解析为byte类型
		if err == io.EOF {
			break
		}
		fmt.Println(byte)
		byteValue = append(byteValue, byte)
		i++
	}
	fmt.Println("运行的次数：", i)
	for _, v := range byteValue {
		writeer.WriteByte(v)
	}
	writeer.Flush()
}
