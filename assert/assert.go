package assert

import (
	"fmt"
)

type inter interface{}

type huamn struct {
	name string
	age int
}

type list []inter
func init() {

	list := make(list,3)
	list[0] = 1
	list[1] = huamn{"dd", 2}
	list[2] = "ddfdfd"
	for index, element := range list {
		switch value := element.(type) {
			case int //
			case string// 
			case huamn//
			default
		}
	}

}

