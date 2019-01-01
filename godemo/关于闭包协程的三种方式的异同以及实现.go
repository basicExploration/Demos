package main

import (
	"fmt"
	"time"
)

var testValue = [5]int{1, 2, 3, 4, 5}

func main() {
//A:
	for v := range testValue {
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(1e5)
//B:
	for v := range testValue {
		func() {
			fmt.Println(v)
		}()
	}
  //C:
	for v := range testValue {
		go func(x int) {
			fmt.Println(x)
		}(v)
	}

}


/*
4 都是4的原因是 协程是并发，所以有可能当for都执行完了 那些个协程还没正式运行，所以到最后就都是4（for执行速度快得很。）
4
4
4


0 因为是闭包，不过每次for循环都是立即执行，所以每次都能讲本次的值执行出来。
1
2
3

4 这种方式虽然跟第一种类似不过每个协程都有唯一的一个赋值，（）这里面不是赋值了嘛，所以说它打印出来的数据就很正常喽。
0
2
1
3
*/
