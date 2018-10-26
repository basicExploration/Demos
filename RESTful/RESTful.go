package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func number1(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "<h1>恭喜你来到了主页。</h1>")
}
func number2(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "你好尊敬的 %s 先生", ps.ByName("y"))
}

func main() {
	router := httprouter.New()
	router.GET("/", number1)
	router.GET("/index/:y", number2)
	log.Fatal(http.ListenAndServe(":8080", router))
}

// 运算符就是简单的 & 和 * 一个取地址、一个解析地址。
