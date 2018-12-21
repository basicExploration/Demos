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

	"golang.org/x/net/html"
)
var sy sync.WaitGroup
func main() {

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
	for i := 0;i < len(v) ;i++  {
		t := i
		go read(v[t],t)

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
	for c := n.FirstChild; c != nil; c = c.NextSibling {//  这里 是c.nextSibling  不然 没办法往下走。
		links = visit(links, c)
	}
	return links
}


func read(doc string,t int) {
	defer sy.Done()
	s := strconv.FormatInt(int64(t),10)
	file,err:= os.Create("./d/"+s+".html")
	fmt.Println("creat err",err)
	if strings.Index(doc,"https://") == -1  {
		if strings.Index(doc,"http://") == -1 {
			doc = "https://coastroad.net" +doc
		}
	}
	res,err := http.Get(doc)
	fmt.Println(err)
	defer res.Body.Close()
	b,_ := ioutil.ReadAll(res.Body)
	write := bufio.NewWriter(file)
	write.Write(b)
	write.Flush()
}
//todo: 未解决并发 请求过多问题，明天解决。先睡觉