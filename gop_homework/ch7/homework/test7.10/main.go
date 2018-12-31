//练习 7.10： sort.Interface类型也可以适用在其它地方。编写一个IsPalindrome(s sort.Interface) bool
// 函数表明序列s是否是回文序列，换句话说反向排序不会改变这个序列。
// 假设如果!s.Less(i, j) && !s.Less(j, i)则索引i和j上的元素相等。
package main

import (
	"fmt"
	"sort"
)

// 函数应该只使用接口，然后行使基础权利。然后传入动态实际值，这样就可以实施真实值了。
func IsPalindrome(s sort.Interface, i int, j int) bool {
	if !s.Less(i, j) && !s.Less(j, i) {
		fmt.Println("则索引i和j上的元素相等")
		return true

	}
	return false
}
