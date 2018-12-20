package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){

	resp,err := http.Get("http://localhost:9090/search")
	defer resp.Body.Close()
	fmt.Println(err)
	b ,err := ioutil.ReadAll(resp.Body)
	fmt.Println(err)
	fmt.Println(string(b))
}
