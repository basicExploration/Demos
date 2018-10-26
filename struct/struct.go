// struc的解释example
package struct

import "fmt"

type person struct {
	name string
	age int
}
func run(value person, value2 int) (int, int){
	if value.age > value2 {
		return value.age, value2
	}
	return "错误的","决定"
}
func init()  {
	bob := person{"hujiaqi", 12}
	var bob1 person
	bob1.string = "ddd"
	bob1.age = 123
	mary := person{age:1234, name:"thomas"}
	date1, date2 := run(bob, 22)
	date3,date4 := run(bob1, 23)
	date5, date6 := run(mary, 233)
	fmt.Print("内容是 %d", date1)
	fmt.Print("内容是 %d", date2)
	fmt.Print("内容是 %d", date3)
	fmt.Print("内容是 %d", date4)
	fmt.Print("内容是 %d", date5)
	fmt.Print("内容是 %d", date6)
}
