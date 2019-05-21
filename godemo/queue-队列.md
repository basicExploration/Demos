## 使用go实现队列
```go
package main

import (
	"fmt"
	"sync"
)

type A struct {
	lock  sync.Mutex
	value []interface{}
}

func NewA(length int) *A {
	return &A{
		value: make([]interface{}, length),
	}
}
func (a *A) Get(key int) interface{} {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value[key]
}

func (a *A) Set(key int, value interface{}) {
	a.lock.Lock()
	a.value[key] = value
	defer a.lock.Unlock()

}

func (a *A) Delete(key int) {
	a.lock.Lock()
	a.value = append(a.value[:key], a.value[key:]...)
	defer a.lock.Unlock()
}


func main(){
	t := NewA(12)
	t.Set(1,"12")
	fmt.Println(t.Get(1))
}
```
