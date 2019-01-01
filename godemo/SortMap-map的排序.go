package main

import (
	"fmt"
	"sort"
)

func main() {
	var j = 0
	re := make([]string, 1000) // 这里设置1000 是忘了 这个slice到底设置多少
	abc := map[string]string{
		"cook":  "可乐",
		"peisi": "百世可乐",
		"dddd":  "fdsfds",
		"ffs2d": "fds",
		"iowq":  "23",
	}
	for i := range abc {
		re[[]byte(i)[0]] = i
		j++
	}
	sort.Strings(re)
	newRe := re[len(re)-len(abc)-1:]
	fmt.Println(newRe)
}



####-

[ cook dddd ffs2d iowq peisi]
