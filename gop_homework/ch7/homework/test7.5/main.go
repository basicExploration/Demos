//练习 7.5： io包里面的LimitReader函数接收一个io.Reader接口类型的r和字节数n，
// 并且返回另一个从r中读取字节但是当读完n个字节后就表示读到文件结束的Reader。
// 实现这个LimitReader函数：
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{
		r,
		n,
	}
}

type LimitedReader struct {
	R io.Reader // 底层Reader接口
	N int64     // 剩余可读取字节数
}

func (l *LimitedReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, io.EOF //  这个EOF是至关重要的，没有它 很多程序都不知道如何退出，因为它是退出信号（type : error）
	}
	if int64(len(p)) > l.N {
		p = p[:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int64(n)
	return
}

func main() {
	read := strings.NewReader("1243")
	rea := LimitReader(read, 2)
	b, err := ioutil.ReadAll(rea)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
	fmt.Println(age())
}

func age()(n int,err error)  {
	return
}
