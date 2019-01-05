package main

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	i := &IntSet{
		[]uint64{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		},
	}
	fmt.Println("开始测试len函数，正确答案应该是10")
	fmt.Println(i.Len())
	fmt.Println("Copy函数，正确答案应该是&{[]uint64{1,2,3,4,5,6,7,8,9,10,}}")
	fmt.Println(i.Copy())
	fmt.Println("开始测试Remove函数，正确答案应该是{1,3,4,5,6,7,8,9,10}")
	i.Remove(1)
	fmt.Println(i.words)
	fmt.Println("开始测试clear函数，正确答案是:[]")
	i.Clear()
	fmt.Println(i.words)
}

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i := &IntSet{
			[]uint64{
				1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
			},
		}
		i.Len()
		i.Remove(1)
		i.Copy()
		i.Clear()
	}
}
