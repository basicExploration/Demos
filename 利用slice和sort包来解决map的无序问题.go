// 将map的k 赋值给一个slice，然后对这个slice的值进行排序（sort排序的是值不是k，因为slice的k是数字没有必要排序）
// 然后将根据这个map以此变量并且取map得值，因为slice的range是有序的，并且 使用sort的时候其实是赋值给了一个新的slice。
package main

import (
	"fmt"
	"sort"
)

func main() {
	a := make(map[string]string)
	a["tr"] = "23"
	a["dd"] = "dd"
	rangeSliceValue := sortMap(a)
	for _, v := range rangeSliceValue {
		fmt.Println(a[v])
	}

}

func sortMap(ma map[string]string) []string {
	var t []string
	for k, _ := range ma {
		t = append(t, k)
	}
	sort.Strings(t)
	return t
}
