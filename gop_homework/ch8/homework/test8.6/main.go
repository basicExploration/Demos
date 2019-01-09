//练习 8.6： 为并发爬虫增加深度限制。也就是说，如果用户设置了depth=3，那么只有从首页跳转三次以内能够跳到的页面才能被抓取到。
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
	"sync"
)

var sy sync.WaitGroup

func main() {

	d := []string{
		"<div>1<div>2<div>3</div></div></div>",
		"<div>1<div>2<div>3</div></div></div>",
		"<div>1<div>2<div>3</div></div></div>",
	}
	sy.Add(len(d))
	for _, k := range d {

		go find(k)
	}
	sy.Wait()
}

func find(d string) {
	defer sy.Done()
	depth := 0
	r := strings.NewReader(d)
	node, err := html.Parse(r)
	if err != nil {
		fmt.Println(err)
	}
	dd(node, depth)

}

func dd(n *html.Node, depth int) {
	depth++
	fmt.Println(n.Data)
	for i := n.FirstChild; i != nil; i = i.NextSibling {
		dd(i, depth)
	}
	if depth > 3 {
		fmt.Println("结束一个线程。")
		return

	}
}
