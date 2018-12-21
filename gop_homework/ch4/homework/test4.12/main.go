//练习 4.12： 流行的web漫画服务xkcd也提供了JSON接口。例如，一个 https://xkcd.com/571/info.0.json
// 请求将返回一个很多人喜爱的571编号的详细描述。下载每个链接（只下载一次）然后创建一个离线索引。
// 编写一个xkcd工具，使用这些离线索引，打印和命令行输入的检索词相匹配的漫画的URL。
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var content map[string][]byte

func main() {
	fujian := ""
	content := make(map[string][]byte)
	str := new(string)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		*str = request.FormValue("str")
		writer.Write([]byte("dd"))
	})
	http.HandleFunc("/show", func(writer http.ResponseWriter, request *http.Request) {
		if _, ok := content[*str]; !ok {
			url := "https://xkcd.com/" + *str + "/info.0.json"
			res, _ := http.Get(url)
			date, _ := ioutil.ReadAll(res.Body)
			content[*str] = date
			fujian = "没有哦~~~"
		} else {
			fujian = "在本地数据库中有数据哦~"

		}
		copy(content[*str], []byte(fujian))
		writer.Write(content[*str])
	})
	http.HandleFunc("/print", func(writer http.ResponseWriter, request *http.Request) {
		for k := range content {
			fmt.Println(k)
		}
	})
	fmt.Println(http.ListenAndServe(":9090", nil))
}
