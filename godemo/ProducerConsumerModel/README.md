## 生产者和消费者模式

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	x(c())
	
}


func c() chan int{
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i // 生产者
		}
	}()
	return ch
}

func x(c chan int){

	Loop:
		for { //消费者。
			select { // 所以最好的方式就是直接使用select然后不close这样就完美了，使用了select close了，然后select中在不ok判断就会一直
			// 输出一个chan后面类型的零值。 如果使用range就必须close，但是你不知道万一有些不知道什么时候要close的时候呢？

			// 所以综上 应该使用 select然后不close。也不用再selec中判断ok。这是最好的选择
			case h := <-c:
				fmt.Println(h)
			case <-time.After(time.Second):
				break Loop

			}
		}
		fmt.Println("生产者消费者模式搞定")

}
```
