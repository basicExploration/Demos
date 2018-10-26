//method的继承案例
package inheritMethod

import (
	"fmt"
	
)

type human struct {
	height int
	name   string
	age    int
}

type students struct {
	human
	class int
	teacherName string
}

type teacher struct {
	human
	class int
	studentName string
}

func (h *human) sayHi(value string) bool {//这里使用的是指针。
	fmt.Print("这个人的名字是%s,今年是%c岁了，职业是%s", h.name, h.age, value)
	return true
}

func init() {
	xiaoming := students{human{110, "xm", 12}, 1, "王娟"}
	wangjuan := teacher{human{170, "wj", 27}, 1, "xiaoming"} 
	xiaoming.sayHi("学生")
	//这里就是集成，因为xiaoming和wangjuan并没有sayhi这个method却可以引用，原因就是因为，human属于
	// teacher 和 stuents 所以后两者就可以集成包含在huan中的方法。
	wangjuan.sayHi("老师")
}

