package main

import "fmt"

func main() {
	mapFunc(name, 12, 12, 33, 23)
}

type fun func(el int) int

func name(el int) int {// 符合func 其实就是 函数头符合 func(int)int即可。
	return el * 10
}
func mapFunc(fn fun, el ...int) {
	for _, v := range el {
		fmt.Println(fn(v))
	}

}
