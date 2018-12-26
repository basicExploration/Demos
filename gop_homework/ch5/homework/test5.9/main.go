//练习 5.9： 编写函数expand，将s中的"foo"替换为f("foo")的返回值。
package main

import "fmt"

func expand(s string, f func(string) string) string {
	return f(s)
}

func main() {
	s := expand("12", func(s string) string {
		s += "."
		return s
	})
	fmt.Println(s)
}
