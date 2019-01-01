package main

import "os/exec"

func main() {
	cad := exec.Command("grep","hello")// 这个其实是创建一个命令
	cad.Start()
	cad.Wait()
}
