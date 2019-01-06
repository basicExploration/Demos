//练习 7.16： 编写一个基于web的计算器程序。
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type result struct {
	One    int
	Two    int
	Fuhao  string
	Dengyu string
	Resu   int
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
		if r.Method == "GET" {
			tem, _ := template.ParseFiles("./main.html")
			tem.Execute(w, nil)
		}

		if r.Method == "POST" {
			var d  = new(string)
			r.ParseForm()
			r.Body.Read([]byte(*d))
			fmt.Println("显示数据",*d)
			one := r.FormValue("one")
			o, _ := strconv.Atoi(one)
			two := r.FormValue("two")
			t, _ := strconv.Atoi(two)
			huhao := r.FormValue("fuhao")
			var res = result{
				o, t, huhao, "=", 0,
			}
			switch huhao {
			case "jia":
				res.Resu = o + t
				res.Fuhao = "+"
			case "jian":
				res.Resu = o - t
				res.Fuhao = "-"
			case "cheng":
				res.Resu = o * t
				res.Fuhao = "x"
			case "chu":
				res.Resu = o / t
				res.Fuhao = "÷"
			default:
				res.Resu = 0
			}
			tem, _ := template.ParseFiles("./main.html")
			tem.Execute(w, res)
		}
	})
	fmt.Println(http.ListenAndServe(":8090", nil))
}
