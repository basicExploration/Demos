//练习 8.7： 完成一个并发程序来创建一个线上网站的本地镜像，把该站点的所有可达的页面都抓取到本地硬盘。
// 为了省事，我们这里可以只取出现在该域下的所有页面(比如golang.org结尾，译注：外链的应该就不算了。)
// 当然了，出现在页面里的链接你也需要进行一些处理，使其能够在你的镜像站点上进行跳转，而不是指向原始的链接。
package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	sy sync.WaitGroup
)

func main() {
	re, err := http.Get("http://www.haust.edu.cn")
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
				m = append(m, v.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		m = find(m, c) //
	}
	return m

}

func findTo(n []string) {
	sy.Add(len(n))
	don := make(chan struct{}, 20)
	for k, v := range n {
		go func() {
			defer func() {
				if e := recover(); e != nil {
					fmt.Println(e)
				}
			}()
			two(don, v,k) // 为了处理 panic
		}()
	}
}

func two(don chan struct{}, v string,k int) {
	don <- struct{}{} // 加入数据
	defer sy.Done()
	if !strings.HasPrefix(v,"http://") {
		v = "http://coastroad.net" + v
	}
	resp, err := http.Get(v)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	data, _ := ioutil.ReadAll(resp.Body) // 数据形成。
	s := strconv.FormatInt(int64(k),10)
	file, err := os.Create("./t/pp" + s + ".html")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	w := bufio.NewWriter(file)
	w.Write(data)
	w.Flush()
	<-don // 减少数据
}
