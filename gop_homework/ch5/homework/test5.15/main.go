//练习5.15： 编写类似sum的可变参数函数max和min。考虑不传参时，max和min该如何处理，再编写至少接收1个参数的版本。
package main

import (
	"fmt"
)

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
func max(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	for i := 0; i < len(vals)-1; i++ {
		if vals[i] > vals[i+1] {
			vals[i+1], vals[i] = vals[i], vals[i+1]
		}
	}
	return vals[len(vals)-1]
}

func min(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	for i := 0; i < len(vals)-1; i++ {
		if vals[i] < vals[i+1] {
			vals[i+1], vals[i] = vals[i], vals[i+1]
		}
	}
	return vals[len(vals)-1]
}
func maxLeastOne(vals ...int) int {
	if len(vals) < 1 {
		return 0
	}
	for i := 0; i < len(vals)-1; i++ {
		if vals[i] > vals[i+1] {
			vals[i+1], vals[i] = vals[i], vals[i+1]
		}
	}
	return vals[len(vals)-1]
}

func minLeastOne(vals ...int) int {
	if len(vals) < 1 {
		return 0
	}
	for i := 0; i < len(vals)-1; i++ {
		if vals[i] < vals[i+1] {
			vals[i+1], vals[i] = vals[i], vals[i+1]
		}
	}
	return vals[len(vals)-1]
}

func main() {
	fmt.Println(max(1, 2, 3, 4, 53, 5))
	fmt.Println(min(12, 2334, 221, 2))
	 //var _ io.Writer = (*dd)(nil) // 测试代码
}

//type dd struct{// 测试代码
//
//}
//func(d1 *dd)Write(p []byte)(int,error){
//	return 0,nil
//}