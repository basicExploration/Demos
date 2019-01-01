// range 这个全局的函数的意义就是对于slice和map进行迭代。就类似于 js中的for(let i of a ){

//}
package rang

import (
	"fmt"
)

func run() {
	var a []int
	b := append(a, 1, 2, 3, 4, 5)
	for i, v := range b {
		fmt.Print(i, v)//i 就是key v 就是value 
	}
}