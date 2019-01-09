//练习 8.6： 为并发爬虫增加深度限制。也就是说，如果用户设置了depth=3，那么只有从首页跳转三次以内能够跳到的页面才能被抓取到。
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
	"sync"
)

var sy sync.WaitGroup
var depth int
func main() {

	d := []string{
		"<div>one!<div>two!<div>three!</div></div></div>",
		"<div>one@<div>two@<div>three@</div></div></div>",
		"<div>one#<div>two#<div>three#</div></div></div>",
	}
	sy.Add(len(d))
	for _, k := range d {

		go find(k)
	}
	sy.Wait()
}

func find(d string) {
	defer sy.Done()

	r := strings.NewReader(d)
	node, err := html.Parse(r)
	if err != nil {
		fmt.Println(err)
	}
	dd(node)

}

func dd(n *html.Node) {
	if n.Type == html.TextNode {
		depth++
		fmt.Println(n.Data)
		if depth > 0 {
			fmt.Println("结束一个线程。")
			return

		}
	}
	for i := n.FirstChild; i != nil; i = i.NextSibling {
		dd(i)
	}
}
