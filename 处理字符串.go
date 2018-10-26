package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Value struct {
	title   string // 第一个字段
	price   string // 第二个字段
	quality string // 第三字段
}

func main() {
	va := new(Value)
	va.Find()
}

func Path(filename string) string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("获取路径发生错误:", err)
		os.Exit(1)
	}
	return filepath.Join(pwd, filename)

}
func (v1 Value) Find() {
	result := make([]string, 0)
	realResult := make([]Value, 0)
	linResult := make([]string, 4)
	path := Path("test.txt")
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		fmt.Println("打开文件发生错误", err)
	}
	//ReadCSV(path, file)
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //用；来读取分割的数据
		if err == io.EOF {
			break
		}
		result = append(result, str)
	}
	fmt.Println("result:", result)
	for _, v := range result {
		linResult = strings.Split(v, ";")
		v1.title = linResult[0]
		v1.price = linResult[1]
		v1.quality = linResult[2]
		realResult = append(realResult, v1)
	}
	for k, v := range realResult {
		fmt.Println(k, v)
	}
}
func IfErr(err error) {
	if err != nil {
		fmt.Println("发生错误，错误是:", err)
	}

}


//字符串的：
/*
"The ABC of Go";25.5,;500
"Functional Programming with Go";56;280
"Go for It";45.9;356
"The Go Way";55;500
*/

//输出结果：
/*
result: ["The ABC of Go";25.5,;500
 "Functional Programming with Go";56;280
 "Go for It";45.9;356
 "The Go Way";55;500
]
0 {"The ABC of Go" 25.5, 500
}
1 {"Functional Programming with Go" 56 280
}
2 {"Go for It" 45.9 356
}
3 {"The Go Way" 55 500
}

*/
