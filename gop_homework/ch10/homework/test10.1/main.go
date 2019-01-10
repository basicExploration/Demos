//练习 10.1： 扩展jpeg程序，以支持任意图像格式之间的相互转换，
// 使用image.Decode检测支持的格式类型，然后通过flag命令行标志参数选择输出的格式。
package main

// 描述整个系统，1 使用flag获取要得到的编码的量 2 通过调用image.Decode 来获取初始的图像。
import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png" // register PNG decoder
	"io"
	"os"
)

var (
	t string
)

func init() {
	flag.StringVar(&t, "Y", "", "要得到的图片的属性")
	flag.Parse()
}
func main() {

	file, err := os.Open("./gif.gif")
	if err != nil {
		fmt.Println("err:", err)
	}

	if err := toJPEG(file, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, _, err := image.Decode(in)
	if err != nil {
		return err
	}
	switch t {
	case "png", "PNG":
		return png.Encode(out, img)
	case "jpeg", "jpg", "JPEG", "JPG":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "gif", "GIF":
		return gif.Encode(out, img, &gif.Options{NumColors: 95})
	default:
		return fmt.Errorf("输入的图片格式错误:", err)
	}
	return nil
}
