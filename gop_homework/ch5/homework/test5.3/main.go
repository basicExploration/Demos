//练习 5.3： 编写函数输出所有text结点的内容。注意不要访问<script>和<style>元素,因为这些元素对浏览者是不可见的。
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"regexp"
	"strings"
)

func main() {
	data := make([]string,0)
	d := "<p>fdsfds23434343</p><p>tt</p><p>dd</p><p>cc</p><script>-----====</script><style>.a{dfdsfds}</style>"
	re := regexp.MustCompile("<script>.+</script>")
	d = re.ReplaceAllString(d,"")
	re1 := regexp.MustCompile("<style>.+</style>")
	d = re1.ReplaceAllString(d,"")
	read := strings.NewReader(string(d))
	node, err := html.Parse(read)
	if err != nil {
		fmt.Println(err)
	}
	visit(&data, node)
	fmt.Println(data)
}

func visit(doc *[]string, n *html.Node){
	if n.Type == html.TextNode {
		if n.Data != "style" && n.Data != "script" {
			*doc = append(*doc, n.Data)// 这里的*doc就是原来的doc 所以它可以改变原来的值
		}

	}
	for o := n.FirstChild; o != nil; o = o.NextSibling {
		visit(doc, o) // 这里必须有 doc原因是 doc扩容了
	}

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
