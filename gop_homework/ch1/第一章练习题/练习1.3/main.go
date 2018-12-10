//做实验测量潜在低效的版本和使用了strings.Join的版本的运行时间差异。
//1.6节讲解了部分time包，11.4节展示了如何写标准测试程序，以得到系统性的性能评测。
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("测试两种字符串拼接的速度差")
	low := SpeedTest(LowSpeed)
	fast := SpeedTest(FastSpeed)
	fmt.Println("现在输出两次测试的结果，前面是慢速的后面是快速的", low, fast)

}

//LowSpeed 慢速函数
func LowSpeed() {
	var result string
	for i := 0; i < len(os.Args); i++ {
		result += "" + os.Args[i]
	}
	fmt.Println("LowSpeed输出结果:", result)
}

//FastSpeed 快速函数
func FastSpeed() {
	str := make([]string, 0)
	for i := 0; i < len(os.Args); i++ {
		str = append(str, os.Args[i])
	}
	result := strings.Join(str, " ")
	fmt.Println("FastSpeed输出结果:", result)
}

//SpeedTest 测试速度的函数
func SpeedTest(f func()) time.Duration {
	start := time.Now()
	f()
	end := time.Now()
	return end.Sub(start)

}

//现在输出两次测试的结果，前面是慢速的后面是快速的 25.489µs 5.266µs

//可以说差距非常明显，如果是拼接字符串的方式，如果字符串越来越长那么每次都要复制整个的字符串进行拼接，
// 非常麻烦。使用strings.Join就非常方便了。

// 速度快的原因也很简单，就是它每次的拼接都是将后面的内容放到前面的字符串byte类型的数组的后面，
// 所以它不需要每次都把前面的东西再赋值一次，所以就快了
