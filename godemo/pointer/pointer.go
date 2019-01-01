package pointer

type huamn struct {
	huamn int
	age int
}
func run() {
	//就说明了，通过直接操作指针，可以直接操作变量指向的那唯一的内容
	a := huamn{12, 2}
	p := &a
	//当然你也可以直接使用：
	b := &huamn{12,34}
	//效果是一样的。
	p.huamn = 2222
	fmt.Print(a)
}


func run1() {
	//在a传递给p时候，传递是值，已经是另一个东西了，所以当p改变后 a病没有什么变化
	a := huamn{12, 2}
	p := a
	p.huamn = 2222
	fmt.Print(a)
}



