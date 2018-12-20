package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp,err := http.Get("http://localhost:9090/search")
	defer resp.Body.Close()
	fmt.Println(err)
	b ,err := ioutil.ReadAll(resp.Body)
	fmt.Println(err)
	fmt.Println(string(b))
	fmt.Println("-------")
	fmt.Println(tt1())
	fmt.Println("-------")
	fmt.Println(tt2())
	fmt.Println("-------")
	fmt.Println(tt3())
	fmt.Println("------")
	t := tt4()
	fmt.Println(t())
	fmt.Println(t())
	fmt.Println("-------")
	tt5()
}

//多说一句，defer 那个东西只是最后再执行但是初始化的时候它可是顺序初始化的，所以也就是说
// 0 1
func tt1() int {
	var i = 0
	defer fmt.Println("defer  tt1", i)
	i++
	return i
}

// 0 1
func tt2() int {
	var i = 0
	defer func(i int) {
		fmt.Println("defer tt2:", i)
	}(i)
	i++
	return i
}

// 1 1
func tt3()int{
	var i =0
	defer func() {
		fmt.Println("defer tt3",i)
		i++
	}()
	i++
	return i
}

//1 2
//2
func tt4() func() int {
	var i = 0
	defer func() {
		fmt.Println("defer tt4:",i)
		i++
	}()

	i++
	return func() int {

		return i
	}
}

func tt5(){

	//t := make(map[string]*int)
	//t["12"] = 12
	//fmt.Println(&t[12])// Cannot take the address of '[12]
	b := make([]int,12)
	b[1] = 12
	fmt.Println(&b[1])//0xc00001e128 slice 可以
}

//-------
//defer  tt1 0
//1
//-------
//defer tt2: 0
//1
//-------
//defer tt3 1
//1
//------
//defer tt4: 1
//2
//2
//-------

//预测正确。

// 不管是什么 defer也好 go也好 golang的编译器都是会进行参数初始化的（但是并没有执行函数的那个一步）所以在tt1中
// i已经被初始化了，此时i是0 那么 即便是后来 i++了也无济于事，

//但是在tt中 defer后面是匿名函数，那么它就无法初始化，所以此时的i就是最后的i（因为闭包自动使用引用变量，换句话说闭包一直可以使用那一个变量，不会重新初始化变量）所以此时
//它里面i就没有任何的值，直到已经返回了i以后 它才执行，那么它的i就是1喽。

// 闭包和循环变量都是引用变量。 其实这么说有点错误 真实的是这样的

//1 首先 闭包 它可以一直使用同一个变量环境，也就是说它不会讲那个上层的变量丢弃 一直都是使用的是它，所以你可以看错是引用变量，

// 对于循环体来说 也是一样的，同样的一个变量在增加 每次的一个循环中 和下一次的循环中都是同样一个变量并没有重新赋值或者是重新指向
//新的空间，那么也可以说是所谓的引用变量喽。

// 所以go 初始化 只是初始化的参数 这个函数体内部的东西它是不会初始化的，不然那就跟执行有啥区别？所以不会。

// 同样还有一点 defer的匿名函数的执行 tt中 是不是最后还有个 i++ 但是 并不会影响 返回值 是因为

// 返回的过程是这样的 首先 先运行 return后面的内容 然后直接返回 然后 开始执行defer 然后执行完毕后 关闭函数 所以函数的最后有一个隐藏式的os.exit
//这个exit是执行完毕了任何东西再关闭的。
