package model

import (
	"fmt"
	"log"
	"runtime"
)

// Where 函数 查找错误地址的函数,标准化在ifelse中使用
func Where(err interface{}) {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("发生错误的文件名是%s,它的错误在这%d一行", file, line)
	fmt.Println("打印的错误信息是:", err)
}

// WherePlus 代替了if err != nil 直接在err后直接调用即可。
// 没有返回值 return就省略了就行，因为函数运行完就会自动释放了。
func WherePlus(err interface{}) {
	if err != nil {
		Where(err)
	}
}
