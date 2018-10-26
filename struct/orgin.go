package struct



import (
"fmt"
)
	

type humans struct {
	name string
	age int
}

func init() {
fmt.Print(humans{"MR", 12})
}