package func

import (
	"fmt"
	)

func run() {//这里就类似于脚本语言了，比如js 是可以使用这种函数传递给变量的形式的。
	a := func() float64 {
		return 3.1425
	}
	fmt.Print(a())
}