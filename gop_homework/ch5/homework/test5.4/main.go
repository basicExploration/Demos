//练习 5.4： 扩展visit函数，使其能够处理其他类型的结点，如images、scripts和style sheets。
package main

import (
	"fmt"
)

func main(){
	s := make([]string,1,10)
	s = []string{
		"1",
	}
	visit(s)
	fmt.Println(s)

}

func visit(s []string){
 s =append(s,"44334")
 fmt.Println(&s)
}