//练习 10.2： 设计一个通用的压缩文件读取框架，用来读取ZIP（archive/zip）和POSIX tar（archive/tar）
// 格式压缩的文档。使用类似上面的注册技术来扩展支持不同的压缩格式，然后根据需要通过匿名导入选择导入要支持的压缩格式的驱动包。
package main

import (
	"archive/zip"
	"archive/tar"
	"fmt"
	"os"
)

func main(){
	t := make([]byte,1000)
	rc,err := zip.OpenReader("./t.zip")
	if err != nil {
		fmt.Println(err)
	}
	defer rc.Close()
	fmt.Println(*rc.File[0])
	fmt.Println(rc.File[0])
	file,err := os.Open("./t.tar")
	if err != nil {
		fmt.Println(err)
	}
	tar.NewReader(file).Read(t)
	fmt.Println(t)
}