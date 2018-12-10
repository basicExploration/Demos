//练习 3.3： 根据高度给每个多边形上色，那样峰值部将是红色(#ff0000)，谷部将是蓝色(#0000ff)。
package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: black; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			n := fmt.Sprintf("%.3f", dy)
			str := []byte(n)
			fmt.Printf("str:",str)
			a, b, c := str[1], str[2], str[2]
			a1, _ := strconv.Atoi(string(a))
			b1, _ := strconv.Atoi(string(b))
			c1, _ := strconv.Atoi(string(c))
			if a1 > 250 {
				a1 = 150
			}
			if b1 > 250 {
				b1 = 150
			}
			if c1 > 250 {
				c1 = 150
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:rgba(%d,%d,%d,1);stroke:purple;stroke-width:1'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, a1, b1, c1)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	if z == 0 {
		log.Fatal("输入数字错误")
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0) x*x +y*y
	if math.IsNaN(r) || math.IsInf(r, 0) {
		return 0
	}
	return math.Sin(r) / r
}
