//练习6.1: 为bit数组实现下面这些方法
//
//func (*IntSet) Len() int      // return the number of elements
//func (*IntSet) Remove(x int)  // remove x from the set
//func (*IntSet) Clear()        // remove all elements from the set
//func (*IntSet) Copy() *IntSet // return a copy of the set
package main

import "fmt"

type IntSet struct {
	words []uint64 // bit 数组。
}

func (i *IntSet) Len() int {
	return len(i.words)
}

func (i *IntSet) Remove(x int) {
	if i.Len() <= x {
		return
	}
	i.words = append(i.words[:x], i.words[x+1:]...)
}
func (i *IntSet) Clear() {
	i.words = make([]uint64, 0)
}
func (i *IntSet) Copy() *IntSet {
	t := *i   // 取得实际值的复制值
	return &t //然后把这个复制的值的地址给返回。
}

func main() {
	var i IntSet
	fmt.Println(i.Copy())
	i.Remove(1)
	fmt.Println(i.Len())
	i.Clear()
}
