// 缓存技术,所谓缓存技术就是一模一样的东西先暂时存一份然后第二次再用的时候直接从内存或者很快的地方直接取就行了的技术
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
		defer res.Body.Close()
		by, err := ioutil.ReadAll(res.Body)
		return by, err
	}
	t := New(f)
	start1 := time.Now()
	t.Get(key)
	end := time.Now()
	fmt.Println("第一次调用的时间", end.Sub(start1), key)

	start2 := time.Now()
	t.Get(key)
	end2 := time.Now()
	fmt.Println("第二次调用的时候", end2.Sub(start2), key)

}

var n = new(sync.WaitGroup)

func main() {
	url := []string{
		"https://www.baidu.com",
		"https://www.taobao.com",
		"https://www.google.com",
		"https://github.com",
	}
	for _, key := range url {
		n.Add(1)
		go func(key string) {
			defer n.Done()
			s(key)
		}(key)// 这里不这样做 你搞进去的key全是最后一个url ，因为for循环很快。并且这里也不是立即执行函数而是开辟新的协程。
	}
	n.Wait()

}
