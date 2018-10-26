//主要是为了测试 switch {}
package switchNil

import "fmt"

func run () {

	switch {//当我们使用空switch时候，我们的意思就要说 switch a（a是boolean类型）然后有三个选项
		// true false 或者是 其实类型的东西，也就是等于 switch true {}
		case 1 > 0:
		fmt.Print("您输出的数据是100") 
		case 1 > 10:
		fmt.Print("您输出的数据是20")
		default:
		fmt.Print("enen")
	}
}

