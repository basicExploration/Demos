package crawler

import (
	"fmt"
	"golang.org/x/net/html"
	"image/jpeg"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var (
	url  string = "https://m.930mh.com/manhua/14396/534008.html?p="
	n    int    = 107
	sym  sync.Mutex
	syw  sync.WaitGroup
	syw1 sync.WaitGroup
	m    = make(map[string]string)
)

// 设计理念 1 找到所有的img然后，获取所有的img的src，然后通过get将src创建到每个不同的文件里（使用jep包的decode和encode）
func Crawler() map[string]string {
	syw.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println(r)
				}
			}()
			s := strconv.FormatInt(int64(i), 10)
			res, err := http.Get(url + s)
			if err != nil {
				fmt.Println(err)
			}
			n, err := html.Parse(res.Body)
			if err != nil {
				fmt.Println(err)
			}
			dear(n, s)
		}(i)
	}
	syw.Wait()
	return m

}

func dear(n *html.Node, dd string) {
	if n.Type == html.ElementNode && n.Data == "img" {
		for _, v := range n.Attr {
			if v.Key == "src" {
				go crearFile(dd, v.Val)
			}

		}
	}
}

func crearFile(k, v string) {
	defer syw.Done()
	res, err := http.Get(v)
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	im, _ := jpeg.Decode(res.Body)
	file, _ := os.Create("./p3/" + k)
	jpeg.Encode(file, im, &jpeg.Options{99})
}
