//练习 1.10： 找一个数据量比较大的网站，用本小节中的程序调研网站的缓存策略，
//对每个URL执行两遍请求，查看两次时间是否有较大的差别，并且每次获取到的响应内容
//是否一致，修改本节中的程序，将响应结果输出，以便于进行对比。
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	ch1 := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine// 并发获取url
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	start = time.Now()
	for _, url := range os.Args[1:] {
		go fetch(url, ch1) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch1) // receive from channel ch 但是接受的时候是单线程，因为要方式没有接受完就结束那种情况
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//ThomasHuke$ go run main.go https://baidu.com https://google.com https://mi.com https://taobao.com https://tmall.com https://github.com https://facebook.com
//Get https://mi.com: x509: certificate has expired or is not yet valid
//2.28s    80685  https://github.com
//3.34s    11261  https://google.com
//3.36s   155649  https://baidu.com
//4.10s   563616  https://facebook.com
//4.18s   247668  https://tmall.com
//6.12s   243688  https://taobao.com
//6.12s elapsed
//0.66s    11319  https://google.com
//0.72s    80685  https://github.com
//1.18s   243688  https://taobao.com
//1.19s   247668  https://tmall.com
//1.54s   561585  https://facebook.com
//1.61s   155776  https://baidu.com
//Get https://mi.com: x509: certificate has expired or is not yet valid
//1.91s elapsed

//事实证明第二次访问比第一次快。估计是因为缓存。
