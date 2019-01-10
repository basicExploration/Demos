//练习 10.3: 从 http://gopl.io/ch1/helloworld?go-get=1 获取内容，
// 查看本书的代码的真实托管的网址（go get请求HTML页面时包含了go-get参数，以区别普通的浏览器请求）。
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	res,err := http.Get("http://gopl.io/ch1/helloworld?go-get=1")
	if err != nil {
		fmt.Println(err)
	}
	data,err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}


