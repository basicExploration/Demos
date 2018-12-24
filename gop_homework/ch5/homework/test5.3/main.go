//练习 5.3： 编写函数输出所有text结点的内容。注意不要访问<script>和<style>元素,因为这些元素对浏览者是不可见的。
package main

import (
	"fmt"
	"golang.org/x/net/html"
)

var i = 0

func main() {
	//res, err := http.Get("https://www.facebook.com")
	//if err != nil {
	//	fmt.Println("打印main函数中的错误err", err)
	//}
	//d, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println("err", err)
	//}
	//read := strings.NewReader(string(d))
	//node, err := html.Parse(read)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//data := visit(nil, node)
	//fmt.Println(data)

fmt.Println(testVisit(9))
}

func visit(doc []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		for _, k := range n.Attr {
			if n.Data == "script" || n.Data == "style" {
				continue
			}
			doc = append(doc, k.Val)
		}
	}
	i++
	for o := n.FirstChild; o != nil; o = o.NextSibling {
		doc = visit(doc, o)
	}
	return doc
}

func testVisit(ii int) int {
	if ii == 0 {
		return 100
	}
	fmt.Println(ii)
	ii = testVisit(ii - 1)

	return ii+1
}
