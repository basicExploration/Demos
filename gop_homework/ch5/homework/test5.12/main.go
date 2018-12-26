//练习5.12： gopl.io/ch5/outline2（5.5节）的startElement和endElement共用了全局变量depth，
// 将它们修改为匿名函数，使其共享outline中的局部变量。
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	depth := 0
	var startElement func(*html.Node, int)
	startElement = func(n *html.Node, depth int) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}
	var endElement func(n *html.Node, depth int)
	endElement = func(n *html.Node, depth int) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	forEachNode(doc, depth, startElement, endElement)
	//!-call

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, depth int, pre, post func(*html.Node, int)) {
	if pre != nil {
		pre(n, depth)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, depth, pre, post)
	}

	if post != nil {
		post(n, depth)
	}
}

//!-forEachNode

//!+startend

//func startElement(n *html.Node) {
//	if n.Type == html.ElementNode {
//		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
//		depth++
//	}
//}
//
//func endElement(n *html.Node) {
//	if n.Type == html.ElementNode {
//		depth--
//		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
//	}
//}

//!-startend
