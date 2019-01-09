//练习 8.9： 编写一个du工具，每隔一段时间将root目录下的目录大小计算并显示出来。
package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	du()
}

func du() {
	for {
		cmd := exec.Command("ls", "-i", "/")
		data, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(data))
		time.Sleep(time.Second * 3)
	}

}
