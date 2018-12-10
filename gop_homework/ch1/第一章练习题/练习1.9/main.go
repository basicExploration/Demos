//练习 1.9： 修改fetch打印出HTTP协议的状态码，可以从resp.Status变量得到该状态码。
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {

		fmt.Println("请输入url")
		log.Fatal("范例: go run main.go https://coastroad.net")
	}
	for _, url := range arg {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		w, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Println("打印一下copy的字节数", w)
		fmt.Println("打印一下 HTTP协议状态码：", resp.Status)
	}
}
