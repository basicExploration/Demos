//使用牛顿法求去平方根
package newTon
import (
	"fmt"
)
func newTown(x float64) float64  {
	x1 := x /2
	var c = 1.0
	
	for {
		
		c = ( x1 + x / x1 ) / 2
		if x1 - c <= 0.00001 {
			return c
		}
		x1 = c
	}
	
}
func init() {
	fmt.Print(newTown(12.3))
}