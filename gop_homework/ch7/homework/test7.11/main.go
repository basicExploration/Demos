//练习 7.11： 增加额外的handler让客服端可以创建，读取，更新和删除数据库记录。
// 例如，一个形如 /update?item=socks&price=6 的请求会更新库存清单里一个货品的价格
// 并且当这个货品不存在或价格无效时返回一个错误值。（注意：这个修改会引入变量同时更新的问题）
package main

import (
	"fmt"
	"net/http"
	"strings"
)

type a func(w http.ResponseWriter, r *http.Request)

func (a1 a) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a1(w, r)
}
func main() {
	var a1 a
	a1 = func(w http.ResponseWriter, r *http.Request) {
		if strings.ToUpper(r.URL.Path) == "/UPDATE"{
			fmt.Println("d")
			item := r.FormValue("item")
			price := r.FormValue("price")
			w.Write([]byte(item + price))
		}
	}
	http.ListenAndServe(":8090", a1)
}
