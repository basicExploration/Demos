//练习 7.4： strings.NewReader函数通过读取一个string参数返回一个满足io.Reader接口类型的值
//(和其它值) 实现一个简单版本的NewReader，并用它来构造一个接收字符串输入的HTML解析器（§5.2）
package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

type File struct {
	s        string
	i        int64 // current reading index
	prevRune int
}

func (r *File) Read(p []byte) (n int, err error) { // 返回写入的字节数和返回错误。
	if r.i >= int64(len(r.s)) { // r.i >= int64(len(r.s))的意思就是读取完了。那么返回0和EOF吧赶紧。
		// 一定要有EOF的判断否则使用其他读取reader的函数就会卡死。
		return 0, io.EOF
	}
	r.prevRune = -1
	n = copy(p, r.s[r.i:])
	r.i += int64(n)
	return
}

func main() {
	b, _ := ioutil.ReadAll(NewReader("12"))
	fmt.Println(b)
}

func NewReader(s string) io.Reader {
	return &File{
		s,
		0,
		-1,
	}
}
