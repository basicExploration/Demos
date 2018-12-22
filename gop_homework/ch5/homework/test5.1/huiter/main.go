package main

// todo 如何解决路由 :id 问题。 解决的办法之一 就是 使用一个自定义的路由，然后使用正则来解析这个表达式。

// 使用重新定义handle 接口的方法来实现这个功能（参考 httprouter ),内部十有八九应该是使用了正则。

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/:user/message", func(h http.ResponseWriter, r *http.Request, p httprouter.Params) {
		h.Write([]byte(p.ByName("user")))
	})
	http.ListenAndServe(":9090",router)
}
