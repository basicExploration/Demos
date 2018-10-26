// 零件仅限铸铁
package main

import (
	"fmt"
	"math"
	"strconv"
)

// 定量
// 零件的材料是灰铸铁硬度为：HB174229
const (
	PI float64 = 3.1415926
	// 年产量：
	PTIME int = 120000
)

// 变量
var (
	// 硬度返回
	Y float64
	//加工的孔的直径
	D float64
	// 左侧电动机名称
	engiL string
	// 右侧电动机名称
	engiR string
	// 切削速度
	v float64
	// 进给量
	jinJiLiang float64
	// 转矩 切削功率 切削力
	F float64
	T float64
	P float64
	// 班制
	ban string
)

func main() {
	// 设定 硬度180 钻孔 φ6
	Y, D = 187, 6
	v, jinJiLiang := QieXueLi(Y, D)
	t1, p1, f1 := QieXueLiang(v, jinJiLiang, D, Y)

	// 设定硬度 180 钻孔 φ4
	Y, D = 187, 4
	v, jinJiLiang = QieXueLi(Y, D)
	t2, p2, f2 := QieXueLiang(v, jinJiLiang, D, Y)
	// fmt.Println("左侧ta右侧tb分别是", (t2 * 10), (t2*10 + t1*14))

	//  确定轴的直径
	dR := 7.3 * math.Sqrt(math.Sqrt(t1*0.01))
	dL := 7.3 * math.Sqrt(math.Sqrt(t2*0.01))
	fmt.Println("主轴的直径为 左侧 右侧", dL, dR)

	// p总

	pL := p2 * 10
	pR := p1*14 + p2*10
	fmt.Println("输入的总的功率左侧和右侧分别是", pL, pR)
	// F总
	fL := f2 * 10
	fR := f2*10 + f1*14
	fmt.Println("输出的总力左侧和右侧分别是", fL, fR)
	ShengChanLv("双班制")
}

// QieXueLi 是选择切削力的函数
func QieXueLi(y float64, d float64) (v, jinJiLiang float64) {
	// fmt.Println("判断输入的数据的类型y d 分别是：", reflect.TypeOf(y), reflect.TypeOf(d))
	if y > 160 && y < 200 {
		if d >= 1 && d <= 6 {
			fmt.Println("确定切削速度和进给量分别是:进给量0.1mm/r,切削速度20m/min")
			return 20, 0.1
		} else if d > 6 && d <= 12 {
			fmt.Println("确定切削速度和进给量分别是：进给量0.2mm/r，切削速度20m/min")
			return 20, 0.2
		} else if d > 12 && d <= 22 {
			fmt.Println("确定切削速度和进给量分别是：进给量0.4mm/r，切削速度20m/min")
			return 20, 0.4
		} else if d > 22 {
			fmt.Println("确定切削速度和进给量分别是：进给量0.8mm/r，切削速度20m/min")
			return 20, 0.8
		}

	} else if y >= 200 && y < 241 {

		if d >= 1 && d <= 6 {
			fmt.Println("确定切削速度和进给量分别是：进给量0.05mm/r，切削速度15m/min")
			return 15, 0.05
		} else if d > 6 && d <= 12 {
			fmt.Println("确定切削速度和进给量分别是：进给量0.1mm/r，切削速度15m/min")
			return 15, 0.1
		} else if d > 12 && d <= 22 {
			fmt.Println("确定切削速度和进给量分别是：进给量0.25mm/r，切削速度15m/min")
			return 15, 0.25
		} else if d > 22 {
			fmt.Println("确定切削速度和进给量分别是：进给量0.4mm/r，切削速度15m/min")
			return 15, 0.4
		}

	} else if y >= 241 && y < 400 {

		if d >= 1 && d <= 6 {
			fmt.Println("确定切削速度和进给量分别是：进给量0.08mm/r，切削速度10m/min")
			return 10, 0.08
		} else if d > 6 && d <= 12 {
			fmt.Println("确定切削速度和进给量分别是：进给量0.15mm/r，切削速度10m/min")
			return 10, 0.15
		} else if d > 12 && d <= 22 {
			fmt.Println("确定切削速度和进给量分别是：进给量0.2mm/r，切削速度10m/min")
			return 10, 0.2
		} else if d > 22 {
			fmt.Println("确定切削速度和进给量分别是：进给量0.3mm/r，切削速度10m/min")
			return 10, 0.3
		}

	} else {
		fmt.Println("您输入的硬度值有问题请您再次输入并查看组合机床简明设计手册page130查看选择用量")
	}

	return 0, 0

}

// QieXueLiang 计算的切削量
func QieXueLiang(v, jinJiLiang, d, y float64) (T, P, F float64) {
	T = 10 * math.Pow(d, 1.9) * math.Pow(jinJiLiang, 0.8) * math.Pow(y, 0.6)
	P = T * v / 9740 / math.Pi / d
	F = 26 * d * math.Pow(jinJiLiang, 0.8) * math.Pow(y, 0.6)
	// 改变格式并且保留小数的位置
	T, e := four(T)
	if e != nil {
		fmt.Println(e)
		return
	}
	P, e = four(P)
	if e != nil {
		fmt.Println(e)
		return

	}
	F, e = four(F)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println("计算的一些参数分别是切削力，切削转矩，切削功率：", F, T, P)
	return
}

// 保留四位小数
func four(v float64) (float64, error) {
	return strconv.ParseFloat(fmt.Sprintf("%.4f", v), 4)
}

func ShengChanLv(ban string) {
	var tk, A, Q, TQ float64
	if ban == "一班制" {
		tk = 2350
		A = 100000 / 1.04
		Q = A / tk
		TQ = 27/27 + 27.5/27.5 + 0.02 + (72.5+100)/8000 + 0.1 + 1.2
		Q2 := 60 / TQ
		fmt.Println(Q)
		fmt.Println("TQ= ", Q2)
		fmt.Println("q1 / q", Q2/Q)
	} else if ban == "双班制" {
		tk = 4600
		A = 100000 / 1.04
		Q = A / tk
		TQ = 27/27 + 27.5/27.5 + 0.02 + (72.5+100)/8000 + 0.1 + 1.2
		Q1 := 60 / TQ
		fmt.Println("Q=", Q)
		fmt.Println("Q1= ", Q1)
		fmt.Println("Q1 / Q", Q1/Q)
	} else {
		fmt.Println("请输入‘一班制’ 或者’双班制‘ ")
		return
	}
}
