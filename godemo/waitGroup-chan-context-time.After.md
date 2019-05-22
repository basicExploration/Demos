## waitGroup-chan-context-time.After.md

```go
package main

import (
	"fmt"
	"context"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	wait := new(sync.WaitGroup)
	ctx, cal := context.WithDeadline(context.Background(), now.Add(time.Second*10))
	ctx = context.WithValue(ctx, "key", "value")
	defer cal()
	wait.Add(1)
	go func() {
		defer wait.Done()
		ctx1, call := context.WithDeadline(ctx, now.Add(time.Second*5))
		defer call()
		ctx1 = context.WithValue(ctx1, "key2", "value2")
		for {
			select {
			case <-ctx.Done():
				fmt.Println("打印数据",ctx.Value("key"))
				fmt.Println("是时候第一层提出了")
				return
			default:
				go func() {
					for {
						select {
						case <-ctx1.Done():
							fmt.Println("内部退出", ctx1.Value("key"), ctx1.Value("key2"))
							return
						default:
							time.Sleep(time.Second)
							fmt.Println("内部的内容》》》》")
						}
					}
				}()
				time.Sleep(time.Second)
				fmt.Println("外部。")
			}
			//select {
			//case <- time.After(time.Second*6):
			//	fmt.Println("强制2秒后退出")
			//	return
			//
			//}
		}
	}()
	wait.Wait()
	fmt.Println("嗯结束了哈，开始干正事了")
	for i := 0; i < 12; i++ {
		fmt.Println(i)
	}
}

```
