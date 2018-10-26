package controller

import (
	"fmt"
	"html/template"
	"os"
	"path"
	"path/filepath"
	"self/text/model"
	"strings"

	"github.com/astaxie/beego"
)

// MainController 注册struct
type MainController struct {
	beego.Controller
}

var r model.Result

//Login 注册 get方法
func (m *MainController) Login() {
	// 获取当前的路径
	fi, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fi1 := path.Join(fi, "/view/layout.html")
	fi2 := path.Join(fi, "/view/login.html")
	tem, err := template.ParseFiles(fi1, fi2)
	if err != nil {
		model.Where(err)
	}

	// var r http.ResponseWriter
	tem.ExecuteTemplate(m.Ctx.ResponseWriter, "layout", nil)

}

//Logon 登陆 get方法
func (m *MainController) Logon() {
	fmt.Println("logon is running")
	fi, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fi1 := path.Join(fi, "/view/layout.html")
	fi2 := path.Join(fi, "/view/logon.html")
	tem, err := template.ParseFiles(fi1, fi2)
	if err != nil {
		model.Where(err)
	}

	tem.ExecuteTemplate(m.Ctx.ResponseWriter, "layout", nil)

}
func (m *MainController) Logonplus() {
	err := m.Ctx.Request.ParseForm()
	fmt.Println("logon is running")
	if err != nil {
		model.Where(err)
	}
	f := r.Find(m.Ctx.Request.FormValue("username"), m.Ctx.Request.FormValue("password"))
	// fmt.Println("这是在longon上的find函数", result)
	if f {
		m.Ctx.Output.Body([]byte("您已经成功登陆"))
	} else {
		m.Ctx.Output.Body([]byte("发生错误或您未注册"))
	}

	// tem.ExecuteTemplate(m.Ctx.ResponseWriter, "layout", result)

}

//Deal 注册处理方法
func (m *MainController) Deal() {
	fmt.Println("Deal is running")
	if strings.ToUpper(m.Ctx.Request.Method) != "POST" {
		m.Ctx.Output.Body([]byte("本方案只用于post方法中"))
		return
	}
	err := m.Ctx.Request.ParseForm()
	if err != nil {
		model.Where(err)
	}
	fmt.Println("username是:", m.Ctx.Request.FormValue("username"))
	fmt.Println("赋值前")
	username := string(m.Ctx.Request.FormValue("username")) // ->>> 逻辑错误 此处无法运行
	passworld := string(m.Ctx.Request.FormValue("password"))
	fmt.Println(username, passworld)
	fmt.Println("r.Add 之前正在运行")
	fmt.Println(r.Password)
	r.Password = passworld
	r.Username = username
	r.Add() // ->> 错误
	fmt.Println("r.add 之后正在运行")
	m.Ctx.Output.Body([]byte("您已经注册请打开登陆界面  ->>   :  /logon"))
}
