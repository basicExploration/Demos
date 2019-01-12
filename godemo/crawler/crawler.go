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
	syO  sync.Once
	syO1 sync.Once
	syO2 sync.Once
)

// 设计理念 1 找到所有的img然后，获取所有的img的src，然后通过get将src创建到每个不同的文件里（使用jep包的decode和encode）
func Crawler() {
	t := make(chan struct{}, 20)
	fmt.Println("正在开始执行")
	syw.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			syO.Do(func() {
				fmt.Println("正在获取数据......")
			})
			t <- struct{}{}
			defer func() {
				if r := recover(); r != nil {
					fmt.Println(r)
				}
			}()
			var res *http.Response
			var err error
			var n *html.Node
			s := strconv.FormatInt(int64(i), 10)
			res, err = http.Get(url + s) // TODO：错误的地方在于每次的请求 太快，如果做到当页面有某某东西了再使用get将东西返回过来呢？
			if err != nil {
				fmt.Println(err)
			}
			n, err = html.Parse(res.Body)
			<-t
			res.Body.Close()
			if err != nil {
				fmt.Println(err)
			}
			dear(n, s)
		}(i)
	}
	syw.Wait()
	fmt.Println("over.")
}

func dear(n *html.Node, dd string) {
	t := make(chan struct{}, 20)
	if n.Type == html.ElementNode && n.Data == "img" {
		for _, v := range n.Attr {
			if v.Key == "src" {
				syO1.Do(func() {
					fmt.Println("正在整合数据......")
				})
				go crearFile(dd, v.Val, t)
			}

		}
	}
	for i := n.FirstChild; i != nil; i = i.NextSibling {
		dear(i, dd)
	}
}

func crearFile(k, v string, t chan struct{}) {
	t <- struct{}{}
	defer func() {
		syw.Done()
	}()
	syO2.Do(func() {
		fmt.Println("正在制作图片......")
	})
	res, err := http.Get(v)
	<-t // 限制访问个数
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	im, err := jpeg.Decode(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	file, err := os.Create("./p3/" + k + ".jpg")
	if err != nil {
		fmt.Println(err)
	}
	err = jpeg.Encode(file, im, &jpeg.Options{99})
	if err != nil {
		fmt.Println(err)
	}
	file.Close()
}
