//练习 5.3： 编写函数输出所有text结点的内容。注意不要访问<script>和<style>元素,因为这些元素对浏览者是不可见的。
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func main() {
	d := "<p>fdsfds23434343</p><p>tt</p><p>dd</p><p>cc</p><script>-----====</script>"
	read := strings.NewReader(string(d))
	node, err := html.Parse(read)
	if err != nil {
		fmt.Println(err)
	}
	data := visit(nil, node)
	fmt.Println(data, len(data))
}

func visit(doc []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		if n.Data != "style" && n.Data != "script" {
			doc = append(doc, n.Data)
		}

	}
	for o := n.FirstChild; o != nil; o = o.NextSibling {
		doc = visit(doc, o) // 这里必须有 doc原因是 doc扩容了
	}
	return doc
}

//func testVisit(ii int) int {
//	if ii == 0 {
//		return 100
//	}
//	fmt.Println(ii)
//	ii = testVisit(ii - 1)
//
//	return ii+1
//}
