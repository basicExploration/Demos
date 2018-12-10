//练习 2.2： 写一个通用的单位转换程序，用类似cf程序的方式从命令行读取参数，
//如果缺省的话则是从标准输入读取参数，然后做类似Celsius和Fahrenheit的单位转换，
//长度单位可以对应英尺和米，重量单位可以对应磅和公斤等。
package main

import (
	"flag"
	"fmt"
	"strconv"
)

var (
	weight, help *string
	long         *string
	t string
)

func init() {
	long = flag.String("long", "m", "长")

	// !!: 这里不能使用 *flag.String 这种方式，原因是 这个时候还没有解析直接获取值的话是没有任何内容的。

	weight = flag.String("weight", "kg", "重量")
	help = flag.String("help", "", "帮助")
	flag.Parse()
	fmt.Println(long)
}

func main() {
	fmt.Println("转换为标准单位")
	Transfor(*long)
	Transfor(*help)
	Transfor(*weight)
}

func Transfor(str string) {
	s := flag.Arg(0)
	st, _ := strconv.Atoi(s)
	stt := float64(st)
	switch str {
	case "m":
		fmt.Println(stt)
	case "cm":
		fmt.Println(float64(stt / 100))
	case "mm":
		fmt.Println(float64(stt / 1000))
	case "kg":
		fmt.Println(stt)
	case "g":
		fmt.Println(float64(stt / 1000))
	case "t":
		fmt.Println(float64(stt * 1000))
	case "help":
		Help()
	}
}

func Help() {
	fmt.Println("示例：go run main.go -long mm 12")
}
