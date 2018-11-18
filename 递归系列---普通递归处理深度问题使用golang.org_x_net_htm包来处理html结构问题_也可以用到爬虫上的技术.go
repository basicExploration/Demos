//递归系列---普通递归处理深度问题使用golang.org_x_net_htm包
//来处理html结构问题_也可以用到爬虫上的技术
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	node, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dataSlice := dealDeep(nil, node)
	fmt.Println(dataSlice)
	for _, k := range dataSlice {
		fmt.Println(k)
		resp, err := http.Get(k)
		if err != nil {
			os.Exit(1)
			fmt.Println(err)
		}
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		resp.Body.Close()
		fmt.Println("数据输出：", string(bytes))
	}
}
func dealDeep(links []string, n *html.Node) []string {
	//fmt.Println("探究一下n是什么？", n.Data)// data就是type类型
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = dealDeep(links, c)
	}
	return links
}
