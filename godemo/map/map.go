//测试
package map

import "fmt"

type huam struct {
	name string
	age int
}

var number1 map[string]huam
func init() {
	number1["dfdfd"] = huam{
			"dfd"
			12
	}
	fmt.Print(number1["dfdfd"])
}

