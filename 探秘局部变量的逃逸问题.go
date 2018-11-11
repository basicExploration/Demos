//所谓局部变量的逃逸，就是指在函数内部定义的变量，然后因为赋值给了一个全局变量，
//所以，导致函数及时return了，这个变量也不会被回收，并且这个变量是定义在堆上的，
//这也是一种对于性能有影响的一种操作方式--- 因为会造成额外的内存分配而造成性能上的损失。
package main

var testT *int

func TestValue() {
	var i *int
	*i = 12
	testT = i // 这里的i就不会因为 TestValue的执行完毕而被回收，也就是所谓的它 逃逸了。
}

func main() {
	TestValue()
}


