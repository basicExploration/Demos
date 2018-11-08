package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("./tt.io")
	if err != nil {
		fmt.Println(err)
	}
	err = os.Rename("./tt.io", "./tt.ploop")
	if err != nil {
		fmt.Println("在改名的时候发生了错误：", err)
	}
	writer := bufio.NewWriter(file)
	for _, url := range os.Args[1:] {
		if url == "" {
			fmt.Println("您尚未输入任何路径，请重试")
			os.Exit(1)
		}
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("获取url发生了错误，错误代码是:", err)
		}
		fmt.Println("打印url:", url)
		fmt.Println("打印状态码：", resp.StatusCode)
		_, err = io.Copy(writer, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Println("读取body的时候发生了错误", err)
		}

	}

}
