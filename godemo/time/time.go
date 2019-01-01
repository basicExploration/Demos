//本包主要是为了switch和time的官方包的示例
package time

import (
	"time"
	"fmt"
)

func showTime() {
ti := time.Now().Weekday()
switch time.Now()
case ti :
fmt.Println("周六")
case ti + 1 :
fmt.Println("周日")
case ti + 2 :
fmt.Println("周一")
case ti + 3 :
fmt.Println("周二")
case ti + 4 :
fmt.Println("周三")
case ti + 5 :
fmt.Println("周四")
case ti + 6 :
fmt.Println("周五")
}