package twoGive

import "fmt"

func run() {

	a := make(map[string]int)
	a["hujiaqi"]  = 1
	a["huhuihui"] = 2
	// b := a["hujiaqi"]
	// delete(a, "hujiaqi")
	_, ok := a["hujiaqi"]//判断是否有这个值，ok是true就是有，否则就是false就是没有并且value为零值也就是nil
	fmt.Print(ok)
}