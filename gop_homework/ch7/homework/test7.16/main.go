//练习 7.16： 编写一个基于web的计算器程序。


// 关注几点 1， 使用前端往后端传输只能用post方法，前端从后端取数据只能用get，前端使用本方案的常规模式只能传过来都是string没办法传过来int数据。
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
)

type result struct {
	One    string `json:"one"`
	Two    string `json:"two"`
	Fuhao  string `json:"fuhao"`
	Dengyu string `json:"dengyu"`
	Resu   int `json:"resu"`
}
var ret  = make(map[string]int)
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
		if r.Method == "GET" {
			tem, _ := template.ParseFiles("./main.html")
			tem.Execute(w, nil)
		}

		if r.Method == "POST" {
			data,_ := ioutil.ReadAll(r.Body)
			res := new(result)
			fmt.Println("看看data中有什么？",string(data))
			json.Unmarshal(data,res)
			fmt.Println(res)
			one,_ := strconv.Atoi(res.One)
			two,_ := strconv.Atoi(res.Two)
			switch res.Fuhao {
			case "jia":

				res.Resu = one + two
				res.Fuhao = "+"
			case "jian":
				res.Resu = one - two
				res.Fuhao = "-"
			case "cheng":
				res.Resu = one * two
				res.Fuhao = "x"
			case "chu":
				res.Resu = one / two
				res.Fuhao = "÷"
			default:
				res.Resu = 0
			}
			ret["name"] = res.Resu
		}
	})
	http.HandleFunc("/s", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w,ret["name"])

	})
	fmt.Println(http.ListenAndServe(":8090", nil))
}
