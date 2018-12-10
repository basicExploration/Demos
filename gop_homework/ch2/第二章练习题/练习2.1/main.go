//练习 2.1： 向tempconv包添加类型、常量和函数用来处理Kelvin绝对温度的转换，
//Kelvin 绝对零度是−273.15°C，Kelvin绝对温度1K和摄氏度1°C的单位间隔是一样的。
package main

import (
	"fmt"
	"os"
	"strconv"
)

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//!-

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%.3f°C", c)
}
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.3f°F", f)
}

func Kelvin(c Fahrenheit) Celsius {
	fmt.Println("输入的是:", c)
	return Celsius(c) + AbsoluteZeroC // 只要底层type一样，就可以直接的转换。不需要再转换为底层对象。
}

func main() {
	arg := os.Args[1:]

	for _, v := range arg {
		str, _ := strconv.ParseFloat(v, 10)
		fmt.Println("单位转换后的结果是:", Kelvin(Fahrenheit(str)))
	}

}

// 只要底层是一样的，那么就可以显示的转换，比如 Fahrenheit和Celsius float64 但是一定要显示的转化 go里 int float可以显示的转化不过会出现丢掉
//小数点的行为。并且一定要判断int的最大值，不然会存在一定的bug.
