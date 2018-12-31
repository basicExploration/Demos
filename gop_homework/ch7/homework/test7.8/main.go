//练习 7.8： 很多图形界面提供了一个有状态的多重排序表格插件：
// 主要的排序键是最近一次点击过列头的列，第二个排序键是第二最近点击过列头的列，
// 等等。定义一个sort.Interface的实现用在这样的表格中。比较这个实现方式和重复使用sort.Stable来排序的方式。
package main

import (
	"fmt"
	"sort"
)

type sortV struct {
	lie  string
	biao string
}

type sortValue []*sortV

// Len is the number of elements in the collection.
func (s sortValue) Len() int {
	var i int
	for range s {
		i++
	}
	return i
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (s sortValue) Less(i, j int) bool {
	if len(s[i].lie) > len(s[j].lie) {

		return true
	}
	return false
}

// Swap swaps the elements with indexes i and j.
func (s sortValue) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	a := sortValue{
		&sortV{"433434343", "32"},
		&sortV{"434433232", "323"},
		&sortV{"3434", "33"},
		&sortV{"433434343", "32"},
		&sortV{"3221324234578654323", "323"},
		&sortV{"3434", "33"},
		&sortV{"433434343", "32"},
		&sortV{"33322432", "323"},
		&sortV{"3434", "33"},
		&sortV{"433434343", "32"},
		&sortV{"334324", "323"},
		&sortV{"3434", "33"},
		&sortV{"433434343", "32"},
		&sortV{"334432", "323"},
		&sortV{"3434", "33"},
	}
	sort.Stable(a)
	for _,v := range a  {
		fmt.Println(v)
	}
}

// 本次 实现的模式是按照lie的len的长度来排序。nice。
