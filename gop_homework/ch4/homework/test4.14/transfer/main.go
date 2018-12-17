package main

import "net/http"

func main(){

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer,request,"https://localhost",http.StatusMovedPermanently)
	})
	http.ListenAndServe(":80",nil)
}
