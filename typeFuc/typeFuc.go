// 测试 函数类型，这种写法就类似编译语言例如js了，不过，稍微复杂一点点
package typeFuc

import (
	"fmt"
)
func runGoing(number1 exampleType, number2 exampleType){
	if number1(110) && number2(999) {
		fmt.Print("t!")
	} else {
		fmt.Print("bad)")
	}
}
type exampleType func (int) bool
func over(value int) bool {
	// 判断是否大于100的函数
	if value > 100 {
		return true
	}
	return false
}
func less(value int) bool {
	//判断是否小于 1000 的函数
	if value < 1000 {
		return true
	}
	return false
}
func init() {
	runGoing(over, less)
}