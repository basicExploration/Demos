// 缓存技术,所谓缓存技术就是一模一样的东西先暂时存一份然后第二次再用的时候直接从内存或者很快的地方直接取就行了的技术
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

// A Memo caches the results of calling a Func.
type Memo struct {
	f     Func
	cache map[string]result
}

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// NOTE: not concurrency-safe!
func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key] // 如果失败前面是零值后面是不ok 的false
	if !ok {                   // 这里是关键 如果不存在这个键值对那么久要赋值
		res.value, res.err = memo.f(key) // 这个地方就是赋值。
		memo.cache[key] = res
	}
	return res.value, res.err
}

func s(key string) {

	f := func(key string) (interface{}, error) {
		res, err := http.Get(key)
		if err != nil {
			fmt.Println("请求资源发生错误,错误代码：", err)
			os.Exit(1)
		}
		defer res.Body.Close()
		by, err := ioutil.ReadAll(res.Body)
		return by, err
	}
	t := New(f)
	start1 := time.Now()
	da, err := t.Get(key)
	dat := da.([]byte)
	fmt.Println(string(dat), err)
	end := time.Now()
	fmt.Println("第一次调用的时间", end.Sub(start1), key)

	start2 := time.Now()
	da, err = t.Get(key)
	fmt.Println(da, err)
	end2 := time.Now()
	fmt.Println("第二次调用的时候", end2.Sub(start2), key)

}

var n = new(sync.WaitGroup)

func main() {
	startt := time.Now()
	url := []string{
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
		"https://www.baidu.com",
		"https://www.mi.com",
		"https://www.mi.com",
		"https://github.com",
	}
	for _, key := range url {
		n.Add(1)
		go func(key string) {
			defer n.Done()
			s(key)
		}(key)// 这里不这样做 你搞进去的key全是最后一个url ，因为for循环很快。并且这里也不是立即执行函数而是开辟新的协程。
		
		
// 			func main() {
// 				for i := 0; i < 10; i++ {
// 					func() {
// 						fmt.Println(i)// 立即执行函数不需要那种写法因为每次循环都是执行完了才能执行下一个。
// 		
// 					}()
// 				}
// 			}

	}
	n.Wait()

}
