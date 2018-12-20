package main

import "fmt"

func main() {
	var a1 a = a{"12"}
	t(a1)

}

// 试验 接口的断言

type a struct {
	value string
}

type ber interface {
	get()
}

type cer interface {
	post()
}

func (a1 a) get() {
	fmt.Println(a1.value)
}

func (a1 a) post() {
	fmt.Println(a1.value + "p")
}

func t(b1 ber){
	type cer interface {// 这就是为了验证 已经实现了ber的变量是否也实现了cer
		post()
	}
	if v, ok := b1.(cer); ok {
		v.post()
	}
	b1.get()

}
