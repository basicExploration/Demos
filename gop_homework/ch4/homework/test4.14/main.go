//练习 4.14： 创建一个web服务器，查询一次GitHub，然后生成BUG报告、里程碑和对应的用户信息。
// exec.Commond().Run() 并不能使用全部的命令 cd命令测试无法使用。。。
package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	gID := new(string)
	HTMLTemplate := []string{
		"1.2.html",
		"2.html",
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		t, err := template.ParseFiles(HTMLTemplate...)
		IsErr(err)
		t.Execute(writer, "https://weibo.com/imgoogege")

	})
	http.HandleFunc("/search", func(writer http.ResponseWriter, request *http.Request) {

		if strings.ToUpper(request.Method) == "POST"{
			*gID = request.FormValue("gID")
			fmt.Println(*gID)
			http.Redirect(writer,request,"/search",304)
		}else if strings.ToUpper(request.Method) == "GET" {
			fmt.Println("https://api.github.com/users/"+*gID)
			res,err := http.Get("https://api.github.com/users/"+*gID)
			IsErr(err)
			da,err := ioutil.ReadAll(res.Body)
			IsErr(err)
			fmt.Fprint(writer,string(da))
			res.Body.Close()
		}


	})
	http.ListenAndServe(":9090", nil)
}

func IsErr(err error) {

	if err != nil {
		fmt.Println(err)
	}
}
