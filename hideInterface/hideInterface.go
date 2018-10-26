package hideInterface

import "fmt"

type Student interface {
	method() string
}

type name struct {
	name string
	age  int
}

func (n name) method() string {
	return "zer"
}

func run(s Student) {
	fmt.Println(s.method())
}

func main() {
	var a name
	a.name = "aa"
	a.age = 12

	run(a)
}
