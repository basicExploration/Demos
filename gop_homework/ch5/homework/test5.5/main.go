//练习 5.5： 实现countWordsAndImages。（参考练习4.9如何分词）
package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	ima, wor int
	sy       sync.WaitGroup
	ma       sync.Map
)

func main() {
	//CountWordsAndImages("https://facebook.com")
	Recover()
	sy.Wait()
	ma.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":--:", value)
		if key == "" || value == "" {
			return false
		}
		return true
	})
	fmt.Println("结束")
}

func Recover() {
	CountWordsAndImagesAsync(urls)
}
func CountWordsAndImagesAsync(url []string) {
	ch := make(chan struct{}, 20)
	sy.Add(len(url))
	for i := 0; i < len(url); i++ {
		go func(i int) {
			defer sy.Done()
			defer func() {// !!!!!!!!!!!!!!!!
				//如果是在外部获取recover可以说压根获取不了，想想也是知道的因为
				// 你并不知道主协程和这个协程到底哪个运行到哪了，
				// 所以要在这个协程中搞定这个panic
				if e := recover(); e != nil {
					fmt.Println("发生了错误，错误信息是：", e)
				}
			}()
			start := time.Now()
			ch <- struct{}{} // 空指针类型
			resp, err := http.Get(url[i])

			if err != nil {
				fmt.Println(err)
				return

			}
			n, err := html.Parse(resp.Body)
			if err != nil {
				fmt.Println("err", err)
			}
			defer resp.Body.Close()
			if err != nil {
				fmt.Println(err)
			}
			wor, im := countWordsAndImagesAsync(n)
			ma.Store(url[i]+"   num", wor)
			ma.Store(url[i]+"   image", im)
			end := time.Now()
			timeS := end.Sub(start)
			ma.Store(url[i]+"花费的时间是：", timeS.String())
			<-ch
		}(i)
	}
}

func countWordsAndImagesAsync(n *html.Node) (int, int) {
	wor, im := countWordsAndImages(n)
	return wor, im
}

//CountWordsAndImages does an HTTP GET request for the HTML
//document url and returns the number of words and images in it.
//func CountWordsAndImages(url string) (words, images int, err error) {
//	//s := "<p>p1</p><a>a1<a>a12</a></a><p>p2</p>"  测试代码
//	//re := strings.NewReader(s)
//
//	resp, err := http.Get(url)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	doc, err := html.Parse(resp.Body) // resp.body已经实现了reader接口直接传进去即可。
//	if err != nil {
//		err = fmt.Errorf("parsing HTML: %s", err)
//		return
//	}
//	words, images = countWordsAndImages(doc)
//	//fmt.Printf("words:%d,images:%d\n", words, images)
//	return
//}

func countWordsAndImages(n *html.Node) (int, int) {
	if n.Type == html.ElementNode && n.Data == "img" {
		ima++
	}
	if n.Type == html.TextNode {
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
