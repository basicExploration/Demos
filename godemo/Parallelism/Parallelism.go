//这个包内详细介绍了 golang的并发和并行。之前golang还只能并发，如今，golang可以实现真正的并行
package Parallelism



import (
	"fmt"
	"runtime"
)

func run(value string) {
	for i := 0; i<=10; i++ {
		runtime.Gosched()//让CPU将时间片段让给别人。下次再执行
		fmt.Println(value)
	}
}

func init() {
	go run("ThomasHuke")
	   run("Hello")
}

// Hello
// ThomasHuke
// Hello
// Hello
// ThomasHuke
// Hello
// ThomasHuke
// Hello
// ThomasHuke
// Hello
// ThomasHuke
// Hello
// ThomasHuke
// Hello
// ThomasHuke
// Hello
// ThomasHuke
// Hello
// ThomasHuke
// Hello
// ThomasHuke
// ThomasHuke