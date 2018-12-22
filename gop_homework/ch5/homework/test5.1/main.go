//练习 5.1： 修改findlinks代码中遍历n.FirstChild链表的部分，将循环调用visit，改成递归调用。
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

var sy sync.WaitGroup
var st = make(chan struct{}, 30)

func main() {
	start := time.Now()
	//t6()
	t5()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
func t5() {
	resp, err := http.Get("http://www.haust.edu.cn")
	if err != nil {
		fmt.Println(err)
	}
	s, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	re := strings.NewReader(string(s))
	doc, err := html.Parse(re) // 将接受的html🌲进行解析。
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	v := visit(nil, doc)
	fmt.Println(v)
	sy.Add(len(v))
	for i := 0; i < len(v); i++ {
		t := i
		go read(v[t], t)
	}
	sy.Wait()
	fmt.Println("j结束了")
}
func t6() {
	var tt int
	resp, err := http.Get("http://www.haust.edu.cn")
	if err != nil {
		fmt.Println(err)
	}
	s, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	re := strings.NewReader(string(s))
	doc, err := html.Parse(re) // 将接受的html🌲进行解析。
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	v := visit(nil, doc)
	fmt.Println(v)
	sy.Add(len(v))
	for i := 0; i < len(v); i++ {
		t := i
		go read(v[t], t)
		tt++
		if tt == 30 {// 这样实现会减慢速度
			tt = 0
			time.Sleep(time.Second)// 因为这里并不知道到底应该设置多少，所以这种方法也可以但是不应该取。
		}

	}
	sy.Wait()
	fmt.Println("j结束了")
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" { // n.type 必须是 element的node节点并且data必须是a属性。
		for _, a := range n.Attr { //
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling { //  这里 是c.nextSibling  不然 没办法往下走。
		links = visit(links, c)
	}
	return links
}

func read(doc string, t int) {
	st <- struct{}{} // 发送数据 发送完 然后 就阻塞次 goruntine 知道缓存里有token的位置
	defer func() {
		if re := recover(); re != nil {
			fmt.Println(re)
		}
	}()
	rea(doc, t)
	<-st // 接受数据，此数据的接受发生在 唤醒沉睡着之前，不论是 发送chan的go gorun 还是接受的 只要没有了数据都是讲 此 goruntine阻塞。
}
func rea(doc string, t int) {
	defer sy.Done()
	s := strconv.FormatInt(int64(t), 10)
	file, err := os.Create("./d/" + s + ".html")
	fmt.Println("creat err", err)
	if strings.Index(doc, "https://") == -1 {
		if strings.Index(doc, "http://") == -1 {
			doc = "http://www.haust.edu.cn" + doc
		}
	}
	res, err := http.Get(doc)
	if err != nil {
		os.Remove("./d/" + s + ".html")
	}
	defer res.Body.Close()
	b, _ := ioutil.ReadAll(res.Body)
	write := bufio.NewWriter(file)
	write.Write(b)
	write.Flush()
}


//实现的想法：使用buffered的channel来解决这个问题。其实也就是使用
