//struct 结构体总结：
//这种对象的方式，在go中是独一无二的，因为在对象中的指针对象无序* 就可以去得到真实值，
//这是仅在struct中可以实现的功能，所以在struct（对象）中尽管定义指针类型，压根不用考虑* 就可以去得到值
//map就不行。

package main

import "fmt"

type aty struct {
	b *b
}
type b struct {
	str string
}

func main() {
	ttt := "d"
	a := map[string]*string{
		"d": &ttt,
	}
	fmt.Println(*a["d"])//d
  fmt.Println(a["d"])//0xc0000601c0
  
  //------ 但是使用对象就没问题
  
  ttttt := &aty{// 即便这里是 &aty，下面取值的时候直接 ttttt 即可  无需使用 （*tttt） 用也可以，因为 这就是对象的语法糖的实质。 
		&b{
			"12",
		},
	}
	fmt.Println(ttttt.b.str)//12 这里没有用 *（tttt.b）就直接可以取得 str
  
  
}
