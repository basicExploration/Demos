//练习 5.5： 实现countWordsAndImages。（参考练习4.9如何分词）
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func main() {

	CountWordsAndImages("https://facebook.com")
}

//CountWordsAndImages does an HTTP GET request for the HTML
//document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}
func countWordsAndImages(n *html.Node) (words, images int) {

}
