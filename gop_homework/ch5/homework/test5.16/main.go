//练习5.16：编写多参数版本的strings.Join。
package main

import (
	"bytes"
	"fmt"
)

func Join(a []string, sep ...string) string {
	var bu bytes.Buffer
	for _, v := range sep {
		bu.WriteString(v) // 往缓存中写入数据
	}
	return bu.String() // 提取出数据。
}
func main() {
	fmt.Println(Join(nil, "12", "33", "323", "233232"))
}
