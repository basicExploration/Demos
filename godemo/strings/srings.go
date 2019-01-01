// 目的是为了测试系统定义的String这个接口，因为是系统定义的，所以并不需要制定出来，系统会在使用fmt.printf时候自动调用string这个
//方法。

package strings

import (
	"fmt"
)

type name struct {
	name string
	age  int
}

func (p name) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

func init() {
	a := name{"dd", 12}
	b := name{"ddddd", 34}
	fmt.Print(a, b)
}
