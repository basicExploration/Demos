//练习 3.4： 参考1.7节Lissajous例子的函数，构造一个web服务器，用于计算函数曲面然后返回SVG数据给客户端。服务器必须设置Content-Type头部：
//
//w.Header().Set("Content-Type", "image/svg+xml")
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
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

func TT(height, width int) string {
	result := make([]string, 0)
	t := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	result = append(result, t)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			n := fmt.Sprintf("%.3f", dy)
			str := []byte(n)
			fmt.Printf("str:", str)
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
			d := fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:rgba(%d,%d,%d,1);stroke:purple;stroke-width:1'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, a1, b1, c1)
			result = append(result, d)
		}
	}
	n := fmt.Sprintf("</svg>")
	result = append(result, n)
	return strings.Join(result, "")
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

func main() {
	t := make([]int,0)
	var height *int
	var width *int
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		heightV, _ := strconv.Atoi(request.FormValue("height"))
		height = &heightV
		wideV, _ := strconv.Atoi(request.FormValue("width"))
		width = &wideV
		fmt.Println("打印第一",*width,*height)
		t = append(t,*width,*height)

	})
	http.HandleFunc("/download", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "img/svg+xml")
		fmt.Println("打印第二：",width,height)
		for i := 0;i< len(t) ;i = i+2  {
			writer.Write([]byte(TT(t[i],t[i+2])))
		}


	})
	http.ListenAndServe(":9090", nil)
}
