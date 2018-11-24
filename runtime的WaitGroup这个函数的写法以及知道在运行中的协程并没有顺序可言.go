/练习一下 runtime的WaitGroup这个函数的写法。并且知道在运行中的协程并没有顺序可言，就像主干道和辅路一样，可以同时也可以辅路比主还快
package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("肯定是最开始运行的")
	total := 0
	ch := make(chan int)
	var sy sync.WaitGroup
	for i := 0; i < 10; i++ {
		sy.Add(1)
		go func(i int) {
			fmt.Println("上层")
			ch <- i
			defer sy.Done()
		}(i)
	}

	go func() {
		fmt.Println("下层")
		sy.Wait() // 除了调度器是0 否则一直都是 阻塞。
		close(ch)
	}()
	fmt.Println("主层")
	for c := range ch {
		fmt.Println(c)
		total += c
	}
	fmt.Println(total)
}
/*
第一次
肯定是最开始运行的
主层// 看 最先运行的惊人是在中间的部分
下层
上层
2
上层
3
上层
4
上层
5
上层
上层
6
1
上层
8
上层
0
上层
上层
9
7
45*/

/*
 第二次

肯定是最开始运行的// 不过最开始运行的一定是这句，因为这句之前没有协程。
上层// 看这个地方 不是主层了，也不是下 而是上。所以协程之间的运行的顺序不一定。
上层
主层
1
0
上层
8
上层
9
下层
上层
2
上层
上层
上层
3
上层
上层
6
5
4
7
45



*/
