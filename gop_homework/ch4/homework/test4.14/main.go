//练习 4.14： 创建一个web服务器，查询一次GitHub，然后生成BUG报告、里程碑和对应的用户信息。
// exec.Commond().Run() 并不能使用全部的命令 cd命令测试无法使用。。。
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
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
		t.Execute(writer, nil)

	})

	http.HandleFunc("/search", func(writer http.ResponseWriter, request *http.Request) {

		if strings.ToUpper(request.Method) == "POST" {
			*gID = request.FormValue("gID")
			fmt.Println(*gID)
			if *gID ==""{
				http.Redirect(writer, request, "/", http.StatusMovedPermanently)
			}
			http.Redirect(writer, request, "/search", http.StatusMovedPermanently)
			// 这个地方原来是因为后面的那个code然后会有不同的方式。测试 301挺好用。
			//由于该代码表示页面地址发生了较长久的改变，故Bing[9]和Google[10]等搜索引擎都推荐使用301重定向，以改变搜索引擎中的实际页面地址。

		} else if strings.ToUpper(request.Method) == "GET" {
			res, err := http.Get("https://api.github.com/users/" + *gID)
			IsErr(err)
			da, err := ioutil.ReadAll(res.Body)
			t := new(Name)
			json.Unmarshal(da, t)
			t.Company = regexp.MustCompile("@").ReplaceAllString(t.Company, "")
			IsErr(err)
			template, err := template.ParseFiles("search.html")
			IsErr(err)
			template.Execute(writer, t)
			res.Body.Close()
		}

	})
	http.ListenAndServe(":80", nil)
}

func IsErr(err error) {

	if err != nil {
		fmt.Println(err)
	}
}
