package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	a := map[string]string{
		"a": "a",
		"b": "b",
	}//数据来源使用map是很爽的选择
	file, _ := os.OpenFile("./test.txt", os.O_WRONLY|os.O_CREATE, 0666)//创建生成文件file
	defer file.Close()
	err := json.NewEncoder(file).Encode(a)
	if err != nil {
		fmt.Println(err)
	}
}
