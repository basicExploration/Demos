package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	bi, _ := exec.LookPath("ls")// 查看 命令的位置
	fmt.Println(bi)
	env := os.Environ()//本机环境变量
	fmt.Println(env)
	arg := []string{"ls", "/usr/bin", "-i"}// 真正的命令
	syscall.Exec(bi, arg, env)// 开始调用
}
