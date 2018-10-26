package main

import "fmt"

// 本页面是为了测试

type Student struct {
	name string
}

type Teacher struct {
	name string
}

func (s Student) method() string {
	return s.name
}

func (t Teacher) method() string {
	return t.name
}

type me interface {
	method() string
}

func run(m me) {

	fmt.Printf(me.method())
}

func main() {
	var a []me

	a[0] = Student{"student"}
	a[1] = Teacher{"teacher"}

	run(a[0])
	run(a[1])

}
