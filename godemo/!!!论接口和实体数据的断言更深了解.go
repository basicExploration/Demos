// 断言里 interface.(interfaceB) 也可以，只不过 b中要把前面的那个总的interface加进去。嵌入进去才可以。
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	tt()
}
func tt() {
	var w io.Writer
	w = os.Stdout // 其实这个并不是真的是接口看接口，目的是为了将这个struct赋值给接口
	rw := w.(io.ReadWriter)
	fmt.Println(rw)
	// success: *os.File has both Read and Write
	////w = new(ByteCounter)
	//rw = w.(io.ReadWriter) // panic: *ByteCounter has no Read method
	var a1 a
	var c cc
	a1 = c
	if t, ok := a1.(b); ok {
		t.get()
		t.writ()
		fmt.Println(t)
	}

}

func writeString(w io.Writer, con string) (int, error) {
	type stringWriter interface { // 这个目的是为了验证传入进来的 io.writer实例是否实现了writeString。又是这种
		// 看似接口对接口实质前面是一个struct实例的那种例子。不论怎么样都是为了看出这个
		// struct 是否实现了那个接口。所以讲接口放前面或者后面都可以。前提是 struct
		// 必须借用一个 interface的身份传入进来。
		writeString(string) (int, error)
	}

	if w, ok := w.(stringWriter); ok {
		return w.writeString(con)
	}
	return w.Write([]byte(con))

}

type a interface {
	get()
}

type b interface {
	a
	writ()
}
type cc struct {
}

func (c cc) get() {
	fmt.Println("cc:get")
}

func (c cc) writ() {
	fmt.Println("cc:writ")

}


// 必须要说 前面是一个 interface 然后传入一个struct 括号里是 另一个interface的用法不是很常见，不过不管 他们俩的顺序怎么办目的都是为了一件事
// 这个interface的方法是否被实现，只不过有两种方式 interface.(Type实体)[最常见] interface.(interface)这两种方式用的时候都
//需要往里面传入一个 继承了前面接口的struct或者其他对象（除了指针  接口对象）然后其实前面那种更适合用在switch中，后面那个是这样用的
// 就是要看继承了前面接口的对象是否也继承了后面的那个接口。然后继承了就使用方法。所以括号里面可以是任何类型包括是接口类型。

