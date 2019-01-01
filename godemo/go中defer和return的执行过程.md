 ## return 的执行过程是这样的:

```go

package main

import "fmt"

var i int

func main() {
	t := t()
	fmt.Println(t)
}

func t() int {
	i = 1
	defer func() {
		i = i + 1
		fmt.Println("neng")
		fmt.Println("i在defer中的值是", i)
	}()
	return i * 10
}



```
 值是；
neng 2 10

1. 执行i =1
2. 执行了return i * 10 这个动作了，此时i是1
3. 想了想 诶 忘了 还有个defer 然后执行了defer，然后 结束了。
也就是说go的return 有两个过程，1 返回数据2 结束任务，defer刚好就是夹在中间，所以return 是比 defer更早的存在。



```go
func t() (i int) {
	i = 1
	defer func() {
		i = i + 1
		fmt.Println("neng")
		fmt.Println("i在defer中的值是", i)
	}()
	return
}


```
值是；
neng 2 2

这种情况，嗯 就是 return没有值 然后执行就是这样的 :
1. i 等于1
2. 因为return后面啥都没有 
3. 执行defer 
4. 想停止了 一看 返回值还有要返回的值呢 就开始 return 此时的i值了，然后返回完 就结束了。
所以 此时 defer中i是2 返回的也是2


```go
func t() (i int) {
	i = 1
	defer func() {
		i = i + 1
		fmt.Println("neng")
		fmt.Println("i在defer中的值是", i)
	}()
	return i * 10
}


```
值是；
neng 11 11

1. 先执行 i =1
2. 再执行 i * 10 然后 
3. 执行了i +1
4. 然后 返回了此时的i。

这种情况可以类比为：

```go
func t() (i int) {
	i = 1
	defer func() {
		i = i + 1
		fmt.Println("neng")
		fmt.Println("i在defer中的值是", i)
	}()
	i = i * 10
	return 
}
返回值：
neng 11 11

```
