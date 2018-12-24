//练习 5.4： 扩展visit函数，使其能够处理其他类型的结点，如images、scripts和style sheets。
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func main() {
	data := make([]string, 0)
	s := "<p>43433433232</p><p>胡家齐这是一个很累大的撞击</p><a href='/123'>32</a><script>fdsfd</script>"
	node, _ := html.Parse(strings.NewReader(s))
	visit(&data, node)
	fmt.Println(data)
}

func visit(s *[]string, n *html.Node) {

	if n.Type == html.TextNode {

			fmt.Println(n.Data)


	}
	for g := n.FirstChild; g != nil; g = g.NextSibling {
		visit(s, g)
	}
}
