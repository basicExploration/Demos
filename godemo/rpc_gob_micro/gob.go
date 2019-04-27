package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

var (
	a = []int{
		1,2,3,4,
	}

)

type tt struct {

	Td *string
}
func main(){
	aa := make([]int,0)
	yy :=  tt{}
	b := bytes.Buffer{}
	enc := gob.NewEncoder(&b)
	dec := gob.NewDecoder(&b)
	enc.Encode(a)
	uio := "1222"
	enc.Encode(&tt{&uio})
	dec.Decode(&aa)
	dec.Decode(&yy)
	fmt.Println(aa)
	fmt.Println(*yy.Td)
}

