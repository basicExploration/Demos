//练习 5.5： 实现countWordsAndImages。（参考练习4.9如何分词）
package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

var ima, wor int

func main() {

	CountWordsAndImages("https://facebook.com")
}

//CountWordsAndImages does an HTTP GET request for the HTML
//document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	s := "<p>p1</p><a>a1<a>a12</a></a><p>p2</p>"
	re := strings.NewReader(s)
	doc, err := html.Parse(re) // resp.body已经实现了reader接口直接传进去即可。
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	fmt.Printf("words:%d,images:%d\n", words, images)
	return
}
func countWordsAndImages(n *html.Node) (int, int) {
	if n.Type == html.ElementNode && n.Data == "img" {
		ima++
	}
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
		sRead := strings.NewReader(n.Data)
		scan := bufio.NewScanner(sRead)
		scan.Split(bufio.ScanWords)
		if scan.Scan() {
			wor++
		}
	}
	for i := n.FirstChild; i != nil; i = i.NextSibling {
		wor, ima = countWordsAndImages(i)
	}
	return wor, ima

}

//p1
//a1
//a12
//p2
//words:4,images:0

//这个 出现的过程表示了 首先查找的节点的规律是 查找p然后看看p有没有子节点，没有接查它的兄弟，然后全部没有会返回nil
//也就是说顺序是
//
// 自己-子女-兄弟。
