//1）在包内部，总是应该从 panic 中 recover：不允许显式的超出包范围的 panic()
//2）向包的调用者返回错误值（而不是 panic）。

// recover panic String的一个练习
package main

import (
	"errors"
	"fmt"
)

type ParseError struct {
	Index int
	Word  string
	Err   error
}

func (p *ParseError) String() string {
	return fmt.Sprintf("包的错误是:%q", p.Word)
}
func Parse() (errValue error) {//"向包的调用者返回错误值（而不是 panic）。" 返回的是err而不是直接panic.
	defer func() {
		if err := recover(); err != nil {
			errValue = fmt.Errorf("返回的是：%v", err)
		}
	}()
	therFunc()
	return
}
func OtherFunc() error {
	panic(&ParseError{12, "222", errors.New("ceshi")})

}
func main() {
	fmt.Println(Parse())
}

 //返回的是：包的错误是:"222"  这个是返回值。
