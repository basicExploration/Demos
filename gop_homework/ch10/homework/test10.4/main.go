//练习 10.4： 创建一个工具，根据命令行指定的参数，报告工作区所有依赖包指定的其它包集合。提示：
// 你需要运行go list命令两次，一次用于初始化包，一次用于所有包。你可能需要用encoding/json（§4.5）包来分析输出的JSON格式的信息。
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os/exec"
	"sync"
	"time"
)

var (
	varS string
	sy      sync.Mutex
	RestMap map[string]int
)

type Result struct {
	Imports []string
}

func init(){
	flag.StringVar(&varS,"mo","github.com/googege","搜索所有的包")
	flag.Parse()
}
func main() {
	result := []string{}
	RestMap = make(map[string]int)
	TT(varS)
	for   {
		start := len(RestMap)
		time.Sleep(time.Second)
		end := len(RestMap)
		if start == end {
			fmt.Println(end)
			break
		}
	}
	for k, _ := range RestMap {
		result = append(result, k)
	}
	fmt.Println(result)

}
func TT(s string) {
	re := new(Result)
	cmd := exec.Command("/usr/local/bin/go", "list", "-e", "-json", s)
	data, err := cmd.Output()
	if err != nil {
		fmt.Print(err)
	}
	json.Unmarshal(data, re)
	for _, v := range re.Imports {
		sy.Lock()
		RestMap[v]++
		sy.Unlock()// 如果在lock前调用unlock那么会发生error错误
		go TT(v)
	}

}

//TODO 可以改成并发的。
