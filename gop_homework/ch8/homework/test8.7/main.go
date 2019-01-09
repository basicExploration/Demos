//练习 8.7： 完成一个并发程序来创建一个线上网站的本地镜像，把该站点的所有可达的页面都抓取到本地硬盘。
// 为了省事，我们这里可以只取出现在该域下的所有页面(比如golang.org结尾，译注：外链的应该就不算了。)
// 当然了，出现在页面里的链接你也需要进行一些处理，使其能够在你的镜像站点上进行跳转，而不是指向原始的链接。
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

var (
	sy sync.WaitGroup
)

func main() {
	re, err := http.Get("https://coastroad.net")
	if err != nil {
		fmt.Println(err)
	}
	n, err := html.Parse(re.Body)
	re.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	result := find(nil, n)
	findTo(result)
	sy.Wait()
}

func find(m []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, v := range n.Attr {
			if v.Key == "href" {
				fmt.Println("测试这个val：",v.Val)
				m = append(m, v.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		m = find(m, n) //
	}
	return m

}

func findTo(n []string) {
	fmt.Println("find to 开始了没")
	sy.Add(len(n))
	don := make(chan struct{}, 20)
	for _, v := range n {
		go func() {
			defer func() {
				if e := recover(); e != nil {
					fmt.Println(e)
				}
			}()
			two(don, v) // 为了处理 panic
		}()
	}
}

func two(don chan struct{}, v string) {
	don <- struct{}{} // 加入数据
	defer sy.Done()
	v = "https://coastroad.net" + v
	resp, err := http.Get(v)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	data, _ := ioutil.ReadAll(resp.Body) // 数据形成。
	file, err := os.Create("./t/pp" + v + ".html")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	file.Write(data)
	<-don // 减少数据
}
