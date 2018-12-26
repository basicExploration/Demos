// 练习 5.7： 完善startElement和endElement函数，使其成为通用的HTML输出器
// 要求：输出注释结点，文本结点以及每个元素的属性（< a href='...'>）
// 使用简略格式输出没有孩子结点的元素（即用<img/>代替<img></img>）
// 编写测试，验证程序输出的格式正确(详见11章)
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

var (
	depth int
)

func main() {
	for _, url := range os.Args[1:] {
		err := outline(url)
		if err != nil {
			fmt.Println("err错误信息:", err)
		}
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {

	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling { // 寻找孩子节点的。
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
	fmt.Println("deep",depth)
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		var s string
		for _, v := range n.Attr {
			s += fmt.Sprintf("%s='%s'", v.Key, v.Val)
			s += " "
		}
		fmt.Printf("%*s<%s %s>\n", depth*2, "", n.Data, s)
		depth++
	}
	if n.Type == html.CommentNode {
		fmt.Printf("%*s%s\n", depth*2, "", n.Data)
	}
	if n.Type == html.TextNode {
		fmt.Printf("%*s%s\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.Data != "input" && n.Data != "img" && n.Data != "br" && n.Data != "hr" {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}

	}
}
