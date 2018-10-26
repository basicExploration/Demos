package map

import (
	"fmt"

)
type huamn struct {
	huamn int
	age string
}

func init() {
	var a = map[string]huamn{//这里出现的human意思是整个这个map是 huamn的集合。
		"humber1": huamn{//这里的huamn是可以省略的。
			huamn:1, age:"hujiaqi"
		},
		"numer2": huamn{
			2, "fdffd"
		},
	}//这里是直接定义并且赋值值的那种方式，所以a不是nil这种方式是允许的。
	// 这里的huamn 

	省略版本

	var a = map[string]huamn{//这里出现的human意思是整个这个map是 huamn的集合。
		"humber1": {//这里的huamn是可以省略的。
			1, "hujiaqi"//这里面才是真正要表达的 struct中的东西，其实这个map js中也有
		},
		"numer2": {
			2, "fdffd"
		},
	}

	a["humber1"] //这样调用就好。
}

// 调用

