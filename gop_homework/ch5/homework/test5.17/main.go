//练习5.17：编写多参数版本的ElementsByTagName，
// 函数接收一个HTML结点树以及任意数量的标签名，
// 返回与这些标签名匹配的所有元素。下面给出了2个例子：
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

var (
	testI   int
	depth   int
	problem = []string{"tto", "cc","ttlpus"}
)

func main() {
	s := "<p id='tto'>ddd<a>aaaaa<span id='ttlpus' placeholder='323'></span></a></p><p id='cc'>cccc</p>"
	read := strings.NewReader(s)
	doc, err := html.Parse(read)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range problem {
		visit(doc, v, start, end)
	}

}
func visit(n *html.Node, pro string, start, end func(*html.Node, string) bool) {
	if er := start(n, pro); !er {
		return
	}
	for i := n.FirstChild; i != nil; i = i.NextSibling {
		visit(i, pro, start, end)
	}
	if er := end(n, pro); !er {
		return
	}
}

func start(n *html.Node, problem string) bool {
	switch n.Type {
	case html.ElementNode:
		var str string
		for _, v := range n.Attr {
			str = fmt.Sprintf("<%s %s='%s'></%s>", n.Data, v.Key, v.Val, n.Data)
		}
		for _, v := range n.Attr {
			if v.Key == "id" {
				if v.Val == problem {
					fmt.Println(str)
					break
				}

			}
		}
		//
		//	fmt.Printf("%*s<%s>%d\n", depth*2, "", n.Data, depth)
		//	depth++
		//case html.TextNode:
		//	fmt.Printf("%*s%s%d\n", depth*2, "", n.Data, depth)
		//case html.CommentNode:
		//	fmt.Println(n.Data)
		//case html.DoctypeNode:
		//	fmt.Println(n.Data)
		//case html.ErrorNode:
		//	fmt.Println(n.Data)
		//case html.DocumentNode:
		//	fmt.Println(n.Data)
	}
	return true
}
func end(n *html.Node, p string) bool {
	if n.Type == html.ElementNode {
		depth--
		//fmt.Printf("%*s</%s>%d\n", depth*2, "", n.Data, depth)
	}
	return true
}
