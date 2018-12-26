//练习 5.8： 修改pre和post函数，使其返回布尔类型的返回值。返回false时，
// 中止forEachNoded的遍历。使用修改后的代码编写ElementByID函数，
// 根据用户输入的id查找第一个拥有该id元素的HTML元素，查找成功后，停止遍历。
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

var (
	testI int
	depth int
)

func main() {
	s := "<p>5<a>a1</a></p><p>6</p><"
	read := strings.NewReader(s)
	doc, err := html.Parse(read)
	if err != nil {
		fmt.Println(err)
	}
	visit(doc, start, end)
}
func visit(n *html.Node, start, end func(n *html.Node)) {
	start(n)
	for i := n.FirstChild; i != nil; i = i.NextSibling {
		visit(i, start, end)
	}
	end(n)
}

func start(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		fmt.Printf("%*s<%s>%d\n", depth*2, "", n.Data,depth)
		depth++
	case html.TextNode:
		fmt.Printf("%*s%s\n%d", depth*2, "", n.Data,depth)
	case html.CommentNode:
		fmt.Println( n.Data)
	case html.DoctypeNode:
		fmt.Println(n.Data)
	case html.ErrorNode:
		fmt.Println(n.Data)
	case html.DocumentNode:
		fmt.Println( n.Data)
	}

}
func end(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n%d", depth*2, "", n.Data,depth)
	}
}
