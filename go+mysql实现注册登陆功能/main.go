package main

import (
	"os"
	"path/filepath"
	"self/text/controller"

	"github.com/astaxie/beego"
)

func main() {
	// 设置静态服务器的访问
	f, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	f += "/static"
	beego.SetStaticPath("/static", f)
	//注册
	beego.Router("/login", &controller.MainController{}, "Get:Login")
	// 登陆
	beego.Router("/logon", &controller.MainController{}, "Get:Logon")
	beego.Router("/logonplus", &controller.MainController{}, "Post:Logonplus")
	// 注册后处理
	beego.Router("/deal", &controller.MainController{}, "*:Deal")
	beego.Run()
}
