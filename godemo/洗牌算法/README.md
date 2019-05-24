## 洗牌算法

```go

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	 r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	 ttt := []int{1,2,3,4,5,6,7,8,9,11,12,13,14,15,16,17,18}

	for _,v :=range r1.Perm(len(ttt)){ r.Perm 就是返回这个len长度的slice的随机的角标排列。
		i :=0
	ttt[i],ttt[v] = ttt[v],ttt[i]
		i++
	}
	fmt.Println(ttt)

}


```
