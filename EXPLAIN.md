# explain something of golang.

**可以不重复造轮子 前提是你会造**
---

1. 闭包

所谓闭包就是指一个函数中的函数，并且这个函数可以调用外部的变量并且无论使用多少次，
都可以一直拥有这个变量不回收，那么这个变量可以称为闭包变量，这个变量也可以看做是
引用变量

2. 循环体变量

在循环体中如果你不是使用的立即执行函数，那么你会发现你到最后的时候执行的时候
发现你所有的值全是通一个值举个例子
```go

package main

import "fmt"
func main(){
	tt()
}
func tt() {
	var i int
	for i := 0; i < 1; i++ {
		dd(i)
	}
	i++
}
func dd(i int) {

	fmt.Print(i)
}
```
这个函数执行的时候 tt中打印出来的是0
原因也是很简单，因为go在初始化的时候因为内部是defer并不会立即执行这个函数
但是 但是 但是要注意一点 deder只是放在return后面执行 但是 它的参数同样是
跟其它函数一样先 初始化 参数的（其它函数也是 先初始化参数然后再执行只是
他们的初始化参数和执行 是顺序的）所以 就造成在for执行完后 这里的i是同样的i
以为初始化的参数 i一直会变，那么到最后初始化完成了 i不就是同样的i了吗，而且
循环体内的变量是它自己本身所以i是会变化的（这也就是循环体引用变量，其实不过是这个变量就是
它自己本身罢了）

只有一种情况可以让i不同那就是立即执行函数 或者你让另一个变量等于这个i那么每次循环
这个新的变量都会初始化一次（它又不是循环体）

3. return 和 defer 的执行顺序

在go里

```go

package main

import "fmt"

func tt() int {
	var i = 0
	defer func() {
		fmt.Println(i)
		i++
		fmt.Println(i)
	}()

	return i
}

func main() {
	tt()
}

```
在这个函数中 执行顺序是这样的 首先先初始化 i  = 0
然后 defer无法初始化（参数变量）因为它没有
然后到了return  然后 直接执行了后面的内容 没错 什么都没有就是i而已
然后 开始了return 直接将值返回了，这时候它没有结束 因为有defer
所以开始执行了defer，然后defer中i是几？嗯 是0 因为这个时候i是0了
然后 打印了0 i再次++ i等于1了，然而并没有什么机会去用这个i了，因为
已经return过了，所以这个i就被收回了，加入return后面是一个闭包，那么这个i
就有用了，它就不会被收回。

然后 这个时候函数就结束了。

看一下 这几种特殊情况


```go
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

		return i// 引用变量。
	}
}
--------
package main

import "fmt"
// 0 1
func tt5() (num int) {
	defer func() {
		fmt.Println("dd", num)
		num++
	}()
	return // 使用这种形式 num始终都是一个，并且改变的是自己所以 返回的值是这个函数域内最后的在os.exit()前面的那个的最后的值，所以i的返回值在defer后运行所以是1

}

func main() {
	fmt.Println(tt5())
	fmt.Println(tt6())
}

func tt6() (num int) {
	defer func(num int) {//// 使用这种形式 defer的函数内部的值已经是num的一个拷贝了，所以它里面怎么改变都不影响外部的return
		fmt.Println("dd", num)
		num++
		num++
		fmt.Println("dd1", num)
	}(num)
	num++
	return

}



```

4. 变量的调用和直接改变参数本身

```go
func tt(){
var i
dd(i)
fmt.Println(i)

}

func dd(i int){
i++
fmt.Println(i)
}

// 1 0

```
原因就是dd(i)这个是代表了 赋值，也就是将var的值赋予了dd的形参
可以看做是 d = i
dd（d） 这个叫做赋值 然后值的拷贝或者是指针的传入以及指针的获取实际值是这个地方的问题

然后还有一种是这样的

```go

func tt(){
var i = 0
{
i = 12
}
fmt.Println(i)
}
```
我之前因为跟赋值搞混所以我总是是用
指针来更更改i其实是错误的理解，因为这个地方的i = 12 压根就没有赋值
这一种说法它不过是更改自己的值而已，就像上面的那个函数即使使用指针，
那么更改指针的实际值的时候也是这么干的，所以i = 12 只是这个参数的在
调整自己的值罢了，它是改变的自己，这里就不牵涉到 是值的拷贝还是引用的拷贝了
因为它压根没有拷贝，仅仅是改变自己罢了。

5. 值的方法和指针的方法

首先 指针的方法和值的方法可以互相调用，因为go会自动帮你
比如指针的方法g（）你使用了值来调用那么go会帮你自动取地址相对的如果是
值的方法 go自动帮你取 *

第二个地方
首先 值上面可以有方法 指针上面也是有方法，我们谈的是 关于对象的方法这点先阐述
因为除了struct（对象）其它类型除了 指针和nil都可以有自己的方法
其它类型不讨论

就是指针的方法的时候，那么go会自动帮你取这个对象的*，因为指针没办法取得对象
里面的value值，只能*去得到，但是go帮你取了，所以你可以使用看似
指针直接去值。

6. 关于实现接口

这个地方go很严格，首先就是接口类型的变量不允许取指针，本来它就是引用类型了
虽然slice这种也是引用类型但是go允许你取它的指针，但是接口类型不允许取（取了也没有意义）

而且对于实现一个接口来说如果你是指针类型实现的接口，那么将变量传递给指针类型时

也必须是指针类型的变量，值同样，不允许自动取&或者*
举个例子

```go
type a interface{

get()
}
type b struct {
c value
}

func(b1 *b)get(){
fmt.Print(b1.c)// go帮你自动取了 *b1这个地方不变，go帮你自动取对象的值，在这个地方也可以。
}
func ddc(a1 a){
a1.get()
}

func main(){
b1 := new(b)
var b2 b
ddc(b1)// 正确
ddc(b2)// 错误❌
}



```
7. 关于slice

slice赋值的时候可以不用指定类型

```
type t int
slice ：= make（[]t,10）

slice = [][]int{
{1}, //不需要使用 t{1}
{2},
{3},
}

```

如果slice里面还是slice或者是map等这种引用类型的话是这么处理的

```go

slice ：= make（[][]int,10）
slice[0] = make([]int,4)
// 或者
slice[1] = []int{
1,2,3,}

```
因为 引用类型不初始化的话 本身就是nil 所以会panic


8. 关于变量的初始化

关于这个地方我也出错过

```
func dd(t *int){
fmt.Println(*t)
}
func main(){
var t *int

dd(t)


}


```
这样就会出错，因为 所有的变量都会初始化（go没有声明 go会自动初始化）
但是t是个指针类型，它的初始化就是nil，所以*nil是错误的，正确的方法是


```go

func dd(t *int){
fmt.Println(*t)
}
func main(){
var t int

dd(&t)


}
```

或者优雅一点

```
func dd(t *int){
fmt.Println(*t)
}
func main(){
t := new(int)

dd(t)


}
```

8. slice 关于他的len和cap

不要超过它的len来查找数据。(而不是cap)只要是超过了len就会报错，虽然没有超过cap
但是它的out of range 错误是根据len来定的。

9. 不要获取map的值的地址

```go
t := make(map[string]string)
t["12"] = 12
fm.println(&t["12"])
```
因为map是动态的，所以它的value的值 的地址不是固定的所以go不允许取得
它的地址。

但是 slice可以。

```go

b := make([]int,12)
	b[1] = 12
	fmt.Println(&b[1])//0xc00001e128 slice 可以

```

10. recover的使用只能在 defer中使用(其它地方调用无效果)

```go

func tt(){
defer func(){
if t := recover;t {
    fmt.Print(t)
}
}
dd()//dd里有panic
}

```

11. 关于接口类型的断言
  - 接口实例.(接口类型)
  - 接口实例.(实际类型)

但是这两个的前面 无一例外都需要传入实际的类型也就是变成了
  - 实际类型的实例.(接口类型)
  - 实际类型的实例.(实例类型)

举个例子
```go
type a struct {
  value string
}

type b interface {
  get()
}
type c interface{
  post()
}

func tt(b1 b){
  // 第一种情况
  if v,ok := b1.(b);ok {// 这个实例 相当于实现了这两个interface
    fmt.println(v.post())
    fmt.Println(v.get())
  }

  // 第二种情况
  if v,ok := b1.(a);ok {
    fmt.Println(v.get())
  }
}
```

```go
type a struct {
	value string
}

type ber interface {
	get()
}

type cer interface {
	post()
}

func (a1 a) get() {
	fmt.Println(a1.value)
}

func (a1 a) post() {
	fmt.Println(a1.value + "p")
}

func t(b1 ber){
  // 这个内部的cer必不可少。
	type cer interface {// 这就是为了验证 已经实现了ber的变量是否也实现了cer
		post()
	}
	if v, ok := b1.(cer); ok {// 这个地方隐藏的说明了  a的实例是满足ber的，不然它这一步就会panic然后它还得满足cer不然还会panic所以这一步直接验证了两次。
		v.post()
	}
	b1.get()

}
```

12. 关于递归，
递归其实就是在执行函数里的函数，直到所有函数都结束了，然后就结束即可。举个例子


```go
func testVisit(ii int) int {
	if ii == 0 {
		return 100
	}
	fmt.Println(ii)
	ii = testVisit(ii - 1)

	return ii+1
}

```
它的执行很明显是从外层的初始栈开始往里执行，然后所有栈执行完毕即可，这里我使用了`ii = testVisit(ii -1)` 目的有两个，1 为了让每下一个的ii都少1，2 就是为了获取上一个栈的返回值，然后每次返回都+1 最后的返回值是109 这也证明了每次返回都是从最上层的栈开始往下调用然后到最下面的然后返回。

13. 关于 channel

 chan 的机制是这样的，当一个没有缓存的（有缓存也是一样只是当缓存满了就一样了）chan，显示导入一个数据，这个时候
 这个发送chan的goruntine就睡眠了（阻塞）然后直到这个chan被接受（只要被接收就行，不管是不是在同样一个goruntine)然后这个数据就被获取了，然后开始唤醒这个chan的发送者的那个goruntine。如果没有后续的数据那么这个chan就应该被关闭了可以人工关闭(close)也可能被系统收回。

看一个例子 这是一个有缓存的，并且利用缓存来限制 http请求数量的操作

```go
var st = make(chan struct{},20)// 将访问的数据限制在20
var sy synv.WaitGroup
func main(){
  dd := []string{"htps://...",",,,,,",",,,"}
  sy.Add(len(dd))
  for _,v := range dd {

      go read(dd)//
  }

sy.Wait()
fmt.Println("执行完毕")
}
func read(st){
  defer sy.Done()
  st <- struct{}// 因为是有缓存的chan所以可以保证一直有20个gorutine是不阻塞的。
  // 只要有一个goruntine不是阻塞的就不会造成死锁
  rea(st)
  <- st
}
func rea(st string){
  res,err := http.Get(st)
}
```

只要有一个goruntine不是阻塞的就不会造成死锁，死锁是程序想退出，但是chan内还有东西，没办法退出，但是又没办法运行，造成了无法结束的窘迫，最终就是各个goruntine都是阻塞然而又不能退出的局面。总之 死锁问题有必要再开一个文件来讨论一下。

14. 关于 type

alias的类型和底层可以转化但是不是隐式是显式。

这里分几个内容
- 一就是

```go
type hand func(http....,http.....)

// 例如

httprouter.handle("/",httprouter.handle)
// 这个时候就是
httprouter.Handle("/",func(http....,http....))
// 即可。
```

这种类型的尤其是在函数的调用的时候  要满足 一个hand类型也是很简单 就是函数满足后面那个样式即可

- type hand string

这种情况也是 函数满足后面的那个 type即可 也就是 是string即可 。

15. 关于引用类型

举个 slice说明一下

```go

func main(){
  t := make([]int,0,10)
  t = []int{
    "12",
  }
  fmt.Println(t)
}

func visit(t []string){
  t = append(s,"1221")
}
```
猜一下 输出的是什么？
是 ["12","1221"]吗？
我本来一直以为是，后来我发现其实不是，我们要先证实一个问题，引用类型并不是指针，它是一个数据结构通常是 一个cap 一个len和一个指针对象。所以它本身也是一个实际的值。当这个地方把t传入visit后，其实是值的复制，然后在visit中，t等于了一个append返回的一个新的slice，那么它就不是指向了原来的那个底层数组了，然后它就和原来的那个t不是同一个底层了，那么它的改变并不会影响之前的那个t。那么什么时候会改变呢，也很简单

```go
func visit(t []string){
  t [1] = "112"
}

```

这就叫做赋值，这并没有去进行值的拷贝或者是指针传递。到这里我么可以说一下了

- 15.2 在go里所有的类型只要是传递数据只有两种模式，1 值的拷贝（包括引用类型它的拷贝只不过是拷贝的它的数据组织）2 指针的拷贝，说白了，指针的拷贝也是值的拷贝，因为它本身也是一种值只不过象征了一种钥匙罢了，所以，除了对数据本身进行直接改变，改变他的数据本身，这种行为可以改变它自己，值的传递的话 统统是有拷贝行为。所以我们以后不能把引用类型看成指针，很不一样。如果是指针的话 那么肯定必须要获取值才能去改变，但是引用类型是go的编译器自动的行为，是不可控的，所以把引用类型的本质搞清楚才行。


16. 关于带（bare return）形参及不带形参的返回值和defer的故事

看两个例子：

> 之前的例子说过了，defer只是执行滞后但是参数记住是参数也就是将形参传入实参的过程其实是同步的并没有什么区别。

```go
//  这个例子中返回值是1
func tt()(t int){
	defer func() {
		t ++
	}()
	return
}

// 这个例子中返回值是0
func tt1()int{
	t := 0
	defer func() {
		t ++
	}()
	return t
}
// 0 这个例子证明了 参数顺序执行化。
func tt1()int{
	t := 0
	defer func(t int) {
		t ++
	}(t)
	return t
}
```

原因也是很简单，首先如果是没有形参的返回值，都是在return后面直接返回的，然后再执行defer然后再执行 os.exit() 但是有形参的就不一样了，它必须返回它形参定义的参数之歌例子中就是t，那么t在哪最后一个出现呢？就是在defer中，所以它的执行过程就变成了，找寻最后出现的t（这里出现在defer中）然后直接执行os.exit() 因为它return后面没有东西，所以它和没有形参的return XXX 很不一样。
那么如果是这样的呢？
```go
func age()(n int,err error){
  return
}
```
它会有什么返回结果呢？答案就是`0 nil` ---- 如果只有return 但是却没有出现n和err那么简单 返回值里不是已经初始化了嘛，那么久返回初始化的结果不就好了嘛所以是 `0 nil`(他们的初始化值)
17. 关于buffered

我们在go的执行中经常使用的一种技巧就是限制go并发的速度，那么这个时候buffered变量就可以实现了它的实现是这样的 ` make(chan xxx,number)` 在get请求中一般我都会这么使用`make(chan struct{},20)` 我们定义了一个新的类型就是 struct{} 这个类型是代表了空，当然你也可以使用bool 都可以 struct{} 类型使用的时候 用 struct{}{} 即可。这就代表了这个chan中最多可以暂存number个数据，这就是所谓的缓存技术，也叫做 buffered数据

18. 关于 recover和并发（多goruntine）

如果是在go的多协程中的panic一定要在这个协程中recover否则在主协程的recover根本无法获取这个panic

```go
go func(i int) {
			defer sy.Done()
			defer func() { // 如果是在外部获取recover可以说压根获取不了，想想也是知道的因为你并不知道主协程和这个协程到底哪个运行到哪了，所以要在这个协程中搞定这个panic
				if e := recover();e != nil {
					fmt.Println(e)
				}
			}()
			start := time.Now()
			resp, err := http.Get(url[i])

			if err != nil {
				fmt.Println(err)
			}
			n, err := html.Parse(resp.Body)
			if err != nil {
				fmt.Println("err",err)
				return
			}
			defer resp.Body.Close()
			if err != nil {
				fmt.Println(err)
			}
			wor, im := countWordsAndImagesAsync(nums, ch, n)
			ma.Store(url[i]+"   num", wor)
			ma.Store(url[i]+"   image", im)
			end := time.Now()
			timeS := end.Sub(start)
			ma.Store(url[i]+"花费的时间是：",timeS.String())
		}(i)
```
19. 关于递归的出栈和进栈
递归都有一个进出栈的过程，那么我们可以利用这个东西来完成一个小的技巧，例如整理div
```go
func a(){
  visit(start,end)
}
func visit(start,end func()){
start()//在进栈时执行的函数
  for {
    visit()
  }
  end()// 在出栈时执行的函数。这样一进一出就可以搞定div的排列。<div> 和 </div>
}
```

[关于递归的出栈和进栈](./关于递归的出栈和进栈.md)

20. 关于 函数内部的函数

```go

func t(){
  var d func()int // 使用这种方式一般都是函数内部有递归，如果不实现 声明一下 函数内部的递归函数将无法运行

  d = func()int{
    //fdffd
  }
  d()

  // 或者
var d = func(){

}
d()

总之，不能使用
func()int{

}
int() 在go语言中这种行为不允许

}
```
关于函数内部的类型 倒是很随意
```go

func t(t1 inter){
  type t struct{
    get()
  }
  if d,ok := t1.(t);ok {
    d.get()
  }
  t1.post()
}

// 这种还是很常见的一种函数内部设立 类型的方式。

```
21. 只有接口和nil不能拥有方法。

```go
invalid receiver type io.Writer (io.Writer is an interface type) // 这是使用了接口的报错。

nil is not a type// 这是使用了nil的报错
```

> ps: 永远不要去取接口的指针，没有丝毫的意义。如果取 slice的指针还有些许的意义(比如在append的时候)但是接口的指针有什么意义？
接口本来就没有实际的意义它本来就是一个抽象的东西。而且它本来也就是引用对象。

22 什么是参数的初始化

```go

package main

import "fmt"
// 0 1
func tt5() (num int) {
	defer func() {
		fmt.Println("dd", num)
		num++
	}()
	return // 使用这种形式 num始终都是一个，并且改变的是自己所以 返回的值是这个函数域内最后的在os.exit()前面的那个的最后的值，所以i的返回值在defer后运行所以是1

}

func main() {
	fmt.Println(tt5())
	fmt.Println(tt6())
}

func tt6() (num int) {
	defer func(num int) {//// 使用这种形式 defer的函数内部的值已经是num的一个拷贝了，所以它里面怎么改变都不影响外部的return
		fmt.Println("dd", num)
		num++
		num++
		fmt.Println("dd1", num)
	}(num)
	num++
	return

}


```

这里就要看了，什么叫初始化，即便是defer或者是匿名函数 无论什么吧，主要是有参数的go都会实现将这个函数的参数给初始化了，然后
函数到真正执行的时候就可以直接使用这个参数，那么为什么上面有两种结果，原因就是这里的num 是值的传递，如果改成指针，就没问题了

```go


// 133
func tt6() (num *int) {
	t := 0
	num = &t
	*num = 0
	defer func(num *int) { //// 使用这种形式 defer的函数内部的值已经是num的一个拷贝了，所以它里面怎么改变都不影响外部的return
		fmt.Println("dd", *num)
		*num++
		*num++
		fmt.Println("dd1", *num)
	}(num)
	*num++
	return

}


```

这种形式跟

```go
// 133
func tt6() (num int) {

	defer func() { //// 使用这种形式 defer的函数内部的值已经是num的一个拷贝了，所以它里面怎么改变都不影响外部的return
		fmt.Println("dd", num)
		num++
		num++
		fmt.Println("dd1", num)
	}()
	num++
	return

}
```

其实是一种意思了。


```go
// 13 1
func tt6() ( int) {
  num := 0
	defer func() { //// 使用这种形式 defer的函数内部的值已经是num的一个拷贝了，所以它里面怎么改变都不影响外部的return
		fmt.Println("dd", num)
		num++
		num++
		fmt.Println("dd1", num)
	}()
	num++
	return num

}
```
像这种正常的函数返回 ，顺序上面说过， 先运行return后面的，然后这个时候就该返回值了（因为人家已经明确说了让你返回你还干啥）然后返回后，开始运行defer 然后defer运行后 就os.exit()就ok了。

23. time.After 的用法

它的作用是 当这个系统没有东西了，然后在设置后几秒后运行，如果其他的case一直有东西，那么它是不会被执行的。
因为在select选择的时候只有在其他的case都没有反应了的时候才会去选择time.After 所以它可以用在 比如什么东西都没有数据了以后
然后按照相应的时候去先取消这个东西，当然还有一种应用场景就是，无论如何就是要5分钟后取消，打死都要
那么可以 使用两个select，有一个select就放一个after和一个default即可。或则，使用 context包的withcanceltimeout 这个函数厉害 无论如何 只要设置的分钟数到了，就能立马取消。
因为cancle withcancle那个函数 如果 不执行canle函数 那么ctx.done 就无法运行，这个时候 cancle的关闭就要在执行这个有ctx的函数之前了。就不能使用defer函数来关闭这个，因为一直有东西运行。

24. 关于 string字符串 []byte 以及[]byte的十六进制表示（以string形式储存）

```go
	// 将string字符串，以unicode编码的形式，找到所有的字符的unicode表示，然后返回位一个数组。
	// [72 101 108 108 111] 就是这个数组（slice）
	src := []byte("Hello")
	// 这个encodeStr 是什么呢？它其实就是把这个数组的所有的数字用16进制表示并且没有加[]而已，而是将这个串变成了字符串的形式储存
	// 就是这个“48656c6c6f” 这个字符串其实还是unicode编码只是 用的16进制并且没有[]罢了，一定不要认为它就是"HELLO"
	encodedStr := hex.EncodeToString(src)
	fmt.Println(src)
	// 48656c6c6f -> 48(4*16+8=72) 65(6*16+5=101) 6c 6c 6f
	fmt.Println(encodedStr)
	byteValue,_ := hex.DecodeString(encodeStr)
	string(byteValue) == "Hello"
```
25. 关于json的一个解析的问题

```json
{"code":0,"data":{"ip":"173.82.115.125","country":"美国","area":"","region":"加利福尼亚","city":"洛杉矶","county":"XX","isp":"XX","country_id":"US","area_id":"","region_id":"US_104","city_id":"US_1018","county_id":"xx","isp_id":"xx"}}

```
```go
type Data struct {
	Data Values `json:"data"`
}
type Values struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

```

定义的时候可以缺少字段，但是，不能跟json字段的格式不符合，举个例子 这里的数据是 在json整个文件下的data对象中，那么你需要两个struct 一个是代表整个的json的数据，第二个struct是代表那个data，你看 那个code和其它字段没有定义吧，没有定义无所谓，但是字段的格式一定要遵守 如果直接把Values传进去就是错误的行为，是不会解析的。

26. 关于go template

第一点 如果你使用`template.Execute` 那么你在那个最外边的layout那个文件里不能使用`{{define "layout"}}{{end}}`
如果你想使用`{{define "layout"}}{{end}}` 那么 你需要`	tem.ExecuteTemplate(w,"layout",nil)`那个中间的变量要用最外边的那个模块
所以最好的就是最外边的那个不用模块，然后使用那个没有模块的就ok了

第二点 如果如果你的子模块就是小的模块很多人称作是母模块 我当他们是小模块子模块，他们里面有变量，那么你肯定是最后使用的是layout这个最外面的文件或者说是模块
那么你就要`{{template "son".}}`   `.` 看到了吗 这个点没有这个点 你在最外面的模块也就是最终使用的时候你发现你的变量压根没有导入，这个就是 变量导入的标志
也就是是 你的子有了 如果不导入 那么这个数据就消失了，我被这个地方坑了几个小时。我的天~~~~。