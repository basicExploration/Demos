//练习 8.2： 实现一个并发FTP服务器。服务器应该解析客户端来的一些命令，比如cd命令来切换目录，
// ls来列出目录内文件，get和send来传输文件，close来关闭连接。
// 你可以用标准的ftp命令来作为客户端，或者也可以自己实现一个。
package main

import (
	"github.com/goftp/server"
)

type MyDriverFactory struct {

}
func(m MyDriverFactory) NewDriver() (server.Driver, error){
	return nil,nil
}
func main() {
	factory := &MyDriverFactory{}
	opts    := &server.ServerOpts{
		Factory: factory,
		Port: 2000,
		Hostname: "127.0.0.1",
	}
	server  := server.NewServer(opts)
	server.ListenAndServe()
}
