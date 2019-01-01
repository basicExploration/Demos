// struct deep example
//中文： 此案例详细讲些了关于struct的具体用法。
package struct

import (
	"fmt"
)
type skills []string// another type
type humme struct {//struct
	name string
	age int
	number int
}
type all struct {//all
	humme
	skills
	number int
}
func run1 (){
	dd := all{humme{name:"hu", age:12, number:229292},number:1222}
	dd.skills = []string{"dddageto"}
	dd.skills = append(dd.stills, "dd", 12, "fewfef")
	fmt.Println("内容就是：", dd.name, dd.age, dd.number, dd.humme.number,dd.skills)
}
