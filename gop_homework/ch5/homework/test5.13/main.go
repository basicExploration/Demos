//练习5.13： 修改crawl，使其能保存发现的页面，必要时，
// 可以创建目录来保存这些页面。只保存来自原始域名下的页面。
// 假设初始页面在golang.org下，就不要保存vimeo.com下的页面。
package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"log"
	"regexp"
	"strings"
)

var (
	url string = "https://coastroad.net"
)

var p map[string]string

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	re := regexp.MustCompile("^http://|^https://")
	for k, v := range list {
		list[k] = strings.Replace(v, url, "", 1)
		if re.MatchString(list[k]) {
			list[k] = ""
		}
	}
	return list
}

func main() {
	p := make(map[string]string)
	for _, v := range crawl(url) {
		if v != "" {
			p[v] = url + v
		}
	}
	fmt.Println(p)
}
