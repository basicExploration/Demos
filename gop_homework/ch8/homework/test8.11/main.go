//练习 8.11： 紧接着8.4.4中的mirroredQuery流程，实现一个并发请求url的fetch的变种。当第一个请求返回时，直接取消其它的请求。
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var sy sync.WaitGroup

func main() {
	mirroredQuery()

}
func mirroredQuery() {
	responses := make(chan string, 3)
	sy.Add(4)
	go func() {
		defer sy.Done()
		responses <- request("https://google.com")
	}()
	go func() {
		defer sy.Done()
		responses <- request("https://facebook.com")
	}()
	go func() {
		defer sy.Done()
		responses <- request("https://youtube.com")
	}()
	ctx, cal := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "message", <-responses)
	cal()// 运行完毕可以取消了。
	go fetch(ctx) // ctx.Done 在cancle（）函数运行后才能运行这个done！！！！！！！
	sy.Wait()
	fmt.Println("over...")

}
func request(hostname string) string {
	res, err := http.Get(hostname)
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(data)
}

func fetch(ctx context.Context) {
	for {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			s := ctx.Value("message").(string)
			fmt.Println("打印第一个s,然后理解取消", s)
			sy.Done()
			return

		default:
			wait()
		}
	}
}

func wait() {
	fmt.Println("waiting...")
}
