//练习 8.10： HTTP请求可能会因http.Request结构体中Cancel channel的关闭而取消。
// 修改8.6节中的web crawler来支持取消http请求。（提示：http.Get并没有提供方便地定制一个请求的方法。
// 你可以用http.NewRequest来取而代之，设置它的Cancel字段，然后用http.DefaultClient.Do(req)来进行这个http请求。）
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "https://coastroad.net", nil)
	if err != nil {
		fmt.Println(err)
	}
	res, _ := http.DefaultClient.Do(req)
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))
	res.Body.Close()

}
