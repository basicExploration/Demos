//练习 1.4： 修改dup2，出现重复的行时打印文件名称。
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	countsArray := make([]map[string]int, 0)
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			if middleCounts := countLines(f, counts); middleCounts != nil {
				countsArray = append(countsArray, middleCounts)
			} else {
				fmt.Println("发生重复错误的文件名字:", f.Name())
			}
			f.Close()
			counts = make(map[string]int)
		}
	}
	fmt.Println("最终的slice是", countsArray)
}

func countLines(f *os.File, counts map[string]int) map[string]int {
	tp := 0
	if strings.Trim(f.Name(), "/dev/stdin") == "" {
		log.Fatal("没有文件，请输入文件")
	}
	input := bufio.NewScanner(f)
	for input.Scan() { // 扫描行.
		counts[input.Text()]++
		tp++
	}
	if len(counts) < tp {
		return nil
	}

	return counts
}

// 总结： 为什么要使用map[string]int 这种形式呢？原因是： counts[input.Text()]++ 这种类型可以自动的统计出来每行是否有重复行，因为
// 一旦是重复行那么最后的counts的len就不等于 for循环的次数，可以说这种方法非常的妙了！


//结果：
//发生重复错误的文件名字: test.txt
//最终的slice是 [map[5:1 6:1 7:1 8:1 1:1 2:1 3:1 4:1]
//map[g:1 k:1 s:1 j:1 l:1 u:1 y:1 h:1 d:1 f:1 tr:1 er:1 we:1 a:1 qe:1]]
