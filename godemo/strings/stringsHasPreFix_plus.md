# 将 strings.index 进行进一步的扩展

```go
func IndexPlus(s, p string){
    for i,value := range s {
        // 使用遍历的方式来遍历字符串的话出来的value结果将是rune unicode类型，所以我们使用string()函数将底层概念一样的字符串改变为string类型。
        b:= strings.HasPrefix(string(value),p)
        fmt.Println("每次输出的结果是:",string(value))
        fmt.Println("是否含有测试数据?",b)
        if b == true {
            fmt.Println("已经找到要找的数据，地址是",i)
        }
    }

}

```

输出结果

```go

func main(){
text("this is the text feeling","f")
}

```

```bash
每次输出的结果是: t
是否含有测试数据? false
每次输出的结果是: h
是否含有测试数据? false
每次输出的结果是: i
是否含有测试数据? false
每次输出的结果是: s
是否含有测试数据? false
每次输出的结果是:  
是否含有测试数据? false
每次输出的结果是: i
是否含有测试数据? false
每次输出的结果是: s
是否含有测试数据? false
每次输出的结果是:  
是否含有测试数据? false
每次输出的结果是: t
是否含有测试数据? false
每次输出的结果是: h
是否含有测试数据? false
每次输出的结果是: e
是否含有测试数据? false
每次输出的结果是:  
是否含有测试数据? false
每次输出的结果是: t
是否含有测试数据? false
每次输出的结果是: e
是否含有测试数据? false
每次输出的结果是: x
是否含有测试数据? false
每次输出的结果是: t
是否含有测试数据? false
每次输出的结果是:  
是否含有测试数据? false
每次输出的结果是: f
是否含有测试数据? true
已经找到要找的数据，地址是 17
每次输出的结果是: e
是否含有测试数据? false
每次输出的结果是: e
是否含有测试数据? false
每次输出的结果是: l
是否含有测试数据? false
每次输出的结果是: i
是否含有测试数据? false
每次输出的结果是: n
是否含有测试数据? false
每次输出的结果是: g
是否含有测试数据? false


```
当然 这样的结果就是可以更加直观的看到每次的输出结果 你还可以将 `strings.HasPreFix`换成 `strings.HasstufFix` 或者是 `strings.Contains` 都是可以的
这样做的目的就是可以将每次的结果更加直观的输出。

或许在某些场景里需要这样的方式呢？对吧。
