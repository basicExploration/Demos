//练习 7.3： 为在gopl.io/ch4/treesort (§4.4)的*tree类型实现一个String方法去展示tree类型的值序列。
package main

import (
	"fmt"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	value := strconv.Itoa(t.value)
	return "这就是我的输出：" + value
}

func Tree() string {
	var t tree
	return fmt.Sprint(&t)
}
