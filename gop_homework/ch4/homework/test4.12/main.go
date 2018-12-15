//练习 4.12： 流行的web漫画服务xkcd也提供了JSON接口。例如，一个 https://xkcd.com/571/info.0.json
// 请求将返回一个很多人喜爱的571编号的详细描述。下载每个链接（只下载一次）然后创建一个离线索引。
// 编写一个xkcd工具，使用这些离线索引，打印和命令行输入的检索词相匹配的漫画的URL。
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://xkcd.com/571/info.0.json")
	fmt.Println(err)
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println(err)
	fmt.Println(string(data))
	res.Body.Close()
}
