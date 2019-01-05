//练习 7.12： 修改/list的handler让它把输出打印成一个HTML的表格而不是文本。html/template包(§4.6)可能会对你有帮助。
package main

import (
	"fmt"
	"go/types"
	"html/template"
	"net/http"
)

type database int
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
			v := `
					<table>
							{{$name,$dd := range .}}
							<tr>$name<td>$dd</td></tr>
							{{end}}
					</table>

				`
			t,_ := template.New("name").Parse(v)
			t.Execute(w,db)
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}
