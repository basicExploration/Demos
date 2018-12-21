package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
    bubbleSort(num)
	end := time.Now()

fmt.Println(end.Sub(start))
	start1 := time.Now()
	bubbleSort1(num)
	end1 := time.Now()

	fmt.Println(end1.Sub(start1))
}
// 冒泡法的意思就是 把最大的或者是 最小的 往后挤压。就可以了 所以 他的 i 圈的循环次数一样。 j圈的循环次数也是一样的。
func bubbleSort(num []int) {
	for i := 0; i < len(num)-1; i++ {
		for j := 0; j < len(num)-i-1; j++ {// 这是 前面大后面小的排列方法
			if num[j]>num[j+1]{
				num[j+1],num[j] = num[j],num[j+1]
			}
		}
	}
}

func bubbleSort1(num []int) {
	for i := 0; i < len(num)-1; i++ {
		for j := 0; j < len(num)-i-1; j++ {// 这是 前面大后面小的排列方法
			if num[j]<num[j+1]{
				num[j+1],num[j] = num[j],num[j+1]
			}
		}
	}
}


//func bubbleSortConcurrent(num []int) {
//
//}


// 这种情况下 不合适 采用并发（行）