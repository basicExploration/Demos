//练习 1.8： 修改fetch这个范例，如果输入的url参数没有 http:// 前缀的话，
// 为这个url加上该前缀。你可能会用到strings.HasPrefix这个函数。
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {

		fmt.Println("请输入url")
		log.Fatal("范例: go run main.go https://coastroad.net")
	}
	for _, url := range arg {
		if !strings.HasPrefix(url, "http://") {
			 url = "http://" + url
		}
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
	}
}
