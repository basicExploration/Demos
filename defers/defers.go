package defers

import (
	"fmt"
)
//规则：后进先出
func run1() {
	defer fmt.Println("world")
	defer fmt.Println("wor")	
	      fmt.Print("hello")
}
// 规则： defer可以读取return后面的那个i值
func run2() (i int) {
	defer func() {
		i++ 
		fmt.Print(i)
	}()
	return 1
} 

// 规则：defer先定义后执行

func run3() {
	i := 0
	defer func(){fmt.Print(i)}()//答案是0, 因为这时候已经将i = 0 了所以是0
	i++
	defer func(){fmt.Print(i)}()//答案是1 ,i已经自增所以。

}
