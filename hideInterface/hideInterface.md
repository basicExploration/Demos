# 关于隐式接口

其实隐式接口说白了，还是接口而已，举个例子好了

```go
type S struct {
    name string
    age int
}
func (s *S)String() string  {
    return fmt.Spingf(s.name, s.age)
}
func main()  {
    s = S{"aa", 12}
    fmt.Println(s)
}


```

接下来我展示fmt的一些源码

```go
type Stringer interface { // 定义一个Stringer接口
	String() string//里面的方法是Sring
}

type S struct {// 定义一个struct
	i int
}

func (s *S) String() string {// 将一个方法绑定给S
	return fmt.Sprintf("%d", s.i)
}

func Print(s Stringer) {//实现这个函数，当然这也是关键的一个函数
	println(s.String())
}

func DynamicPrint(any interface{}) {
	if s, ok := any.(Stringer); ok {//这个其实是一个判定，可有可无，但是如果没有的话会有bug的可能
		Print(s)
	}
}

func main() {
	var s S
	s.i = 123456789
	Print(&s)// 这个就是关键了 为什么将一个struct传进去也可以呢？因为这个s并没有赋值给这个interface啊
    //关键就在于，这个interface里面定义的是一个String，并且这有这个方法，同时你的S类型实现了一个方法，String
    //长得像鸭子就是鸭子 根据鸭子理论，那么相当于你讲一个struct类型的数据赋值给了interface，所以，综上所述。
	DynamicPrint(&s)
}


```

再看另一个例子，其实还是这个隐藏式的接口，也就是鸭子理论。


```go

package main

import "fmt"

type Student interface {
	method() string
}

type name struct {
	name string
	age  int
}

func (n name) method() string {
	fmt.Print("sdd")

}

func run(s Student) {
	fmt.Println(s.method())
}

func main() {
	var a name
	a.name = "aa"
	a.age = 12

	run(a)
}

}

```
它之所以能实现，是因为在interface中定义的方法 无论是方法名还是返回值，都跟name中的方法是一样的，那么根据长得像鸭子的就是鸭子的这个鸭子理论，就算是不显示的定义也是可以的。


如果方法和接口中不一致的话就会报错

```go
cannot use a (type name) as type Student in argument to run:
	name does not implement Student (wrong type for method method)
		have method()
		want method() string
main/main.go11014:24
 go-plusLFUTF-8Go
master2 files

```
其实看到 提示的内容非常的人性化，例如它说 是 method()不过想要的是 method() string

并且错误信息也很明确。 不能使用 name类型的变量传入 student的变量中 ，其实也就是 鸭子理论没有成立。


最后我总结一下，interface 之所以有隐藏式接口，主要原因就是struct和interface可以互换。隐藏式接口无非就是将struct转化为了符合条件的interface罢了。

而且可以说 struct这个东西真的是深入到了go语言中的方方面面，一个struct中可以嵌套另一个stuct并且还可以有两种方式

```go
type name struct {
    name int
    year string
}

type student struct {
    name
    app string
}

========================
也可以 使用

type student struct {
    n name
    app string
}

```
也就是说你可以直接嵌入 也可以使用 一个 类型这种方式 顺便说一下

```go
var a student{pass}

这种形式都是因为 后面的这个东西是一个类型


```


那么，GO语言的interface机制到底是如何实现的呢？

interface value

上述代码中函数Print的参数是一个Stringer接口，也就是Stringer的一个对象实例。这个对象实例叫做interface value。它的数据结构如下：
```go

type iface struct {
    tab \*itab
    data unsafe.Pointer
}

```
其中tab字段类似于C++的vptr，tab中包含了对应的方法数组，除此之外还保存了实现该接口的类型元数据。data是对应的实现该接口的类型的实例指针。

itab数据结构如下：
```go

type itab struct {
    inter     *interfacetype
    _type     *_type
    link      *itab
    bad       int32
    unused    int32
    fun       [0]unsafe.Pointer
}

```
其中inter字段表示这个interface value所属的接口元信息，\_type字段表示具体实现类型的元信息，fun字段表示该interface的方法数组。link,bad,unused字段暂时不关心。

当我们在GO代码中调用一个接口的方法时，操作类似如下： `s.tab->fun[0](s.data)。` 调用开销还是很小的。

Itab的生成方式

一个自定义的结构体可以实现某个接口，然后可以隐式的转换到对应的接口。这种操作有点像C++的派生类转换为基类一样，这个操作是一个运行时绑定过程。而GO语言的interface机制还有一些其他特性：比如一个具体类型可以实现N多方法，但是只有其中某几个或者全部都满足某个接口，而此时，不可能把所有的方法都放到Itab中，这就意味着需要在绑定过程中剔除某些不需要的方法。

GO编译器会在编译时会为每个自定义结构体和interface类型生成一个类型元数据，用来描述这个类型的名称，类型的HASH值，类型的方法列表，方法列表中还包括了方法的名称。而在一个自定义结构体转换到一个interface类型时，GO编译器会生成代码，使其在运行时计算Itab，完成动态绑定方法的需求。这个计算Itab的过程相对来说比较简单，因为GO编译器生成的类型元数据中包含了所有的方法名称和地址，那么在一个结构体实例转换为interface value时，只需要把interface的方法列表作为基，方法名和方法类型作为KEY，去结构体元数据中查找对应的方法即可。

GO的runtime库中对Itab的查找过程做了优化，由O(ni * nt)复杂度变为O(ni + nt)。依据是一个自定义结构体实现的方法一定是大于或等于某个具体interface的方法集的。所以可以事先把所有的方法按照名字从小到大排序，然后在匹配到一个方法后，可以在下次查找时使用上次的索引值。

除此之外，GO编译器为了减少每次不必要的Itab，还增加了一个对应的itab的缓存。你可以编译一个GO程序，然后反编译后可以查看到一个类似go_itab__main_S_main_Stringer名称的变量。在每次一个结构体转换到一个interface之前都会检查这个缓存是否有效，有效就使用。这个检查也只是一个cmp指令而已。

还有在GO运行时库里，为了减少每次的Itab实现，还做了相应的优化。内部实现了一个HASH表，保存了每个具体结构体到interface转换生成的Itab实例。代码可以在`go\src\runtime\iface.go getitab`函数中看到。


想当于 这个interface 的实例它并不是interface了，它是一个struct

也即是说想实现这个interface，那么这个结构体就必须拥有大于等于这个interface中的方法，当然大于的方法，什么方法都可以，在 struct转化为 interface value的时候，会剔除那些没有用的方法。
