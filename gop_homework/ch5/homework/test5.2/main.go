//练习 5.2： 编写函数，记录在HTML树中出现的同名元素的次数。
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	lin := make(map[string]int)
	res, err := http.Get("http://www.haust.edu.cn")
	if err != nil {
		fmt.Println("打印main函数中的错误err", err)
	}
	d, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("err", err)
	}
	read := strings.NewReader(string(d))
	node, err := html.Parse(read)
	if err != nil {
		fmt.Println(err)
	}
	dataMap := visit(lin, node)
	for k,v := range dataMap {
		fmt.Println(k,": ",v)
	}
	res.Body.Close()
}

func visit(s map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {// 出错
		s[n.Data]++
	}
	for i := n.FirstChild; i != nil; i = i.NextSibling {
		s = visit(s, i)
	}
	return s
}
