// 几种 ok 模式
package main

import "fmt"

var a = make(map[string]string)

func main() {
	
	//1 测试 map中的是否存在某个键
	var a1 AT
	var Jet interface{}
	Jet = "12"
	if v,ok := a["12"];ok {
		println(v)
	}
	
	
	// 2 某对象是否实现了接口
	
	if v,ok := a1.(b);ok {
		println(v)
	}
	
	
	//3 空接口的运用
	
	if v,ok := Jet.(int);ok {
		fmt.Println(v)
	}
	
	
	//4 空接口的switch运用

	switch v := Jet.(type) {
	case int:
		fmt.Println(v)
	case string:
		fmt.Println(v)
	default:
		fmt.Println("err")
	}
	
	
	//5 recover 中的运用。
	
	if ok := recover();ok {
		fmt.Println(ok)
	}
	
	
	//6 检测一个通道ch是否关闭 range 方法检测
	
    	for input := range ch {
		...
    	}
	
	
	//7 检测一个通道ch是否关闭，ok方法。 
	
    	for {
        if input, open := <-ch; !open {
            break // 通道是关闭的
	    }
    	}
		
}

type AT interface {
	get()
}

type b struct {
	
}
func(b1 b)get(){
	fmt.Println("1")
}
