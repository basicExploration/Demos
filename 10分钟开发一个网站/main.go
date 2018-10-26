package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/", func(arg1 http.ResponseWriter, arg2 *http.Request) {
		arg1.Header().Add("Content-Type", "text/html")
		arg1.Header().Add("charset", "utf-8")
		// 根据时间戳来进行验证 转换为 Unix时间。
		crutime := time.Now().Unix()
		// 生成md5接口
		h := md5.New()
		// 将格式化好的数据 也就是这个Unix时间戳传递给这个空的接口，使用的是io.WriteString 它的作用就是前者接收后者的值。
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		// 通过 Sprintf来规范输出的格式。
		token := fmt.Sprintf("%x", h.Sum(nil))
		// 解析模板
		r, e := template.ParseFiles("./index.html")
		if e != nil {
			fmt.Println(e)
			plus(arg1, arg2)
			return
		}
		// 生成数据 这里的数据是一个根据时间戳来验证的一个MD5的密匙。
		r.Execute(arg1, token)

	})
	http.HandleFunc("/log", func(arg1 http.ResponseWriter, arg2 *http.Request) {
		// arg1.Header().Add("Content-Type", "text/html")
		arg1.Header().Add("charset", "utf-8")
		arg2.ParseForm()

		if arg2.Method == "GET" {
			t, e := template.ParseFiles("./log.html")
			if e != nil {
				fmt.Println(e)
				plus(arg1, arg2)
				return
			}
			t.Execute(arg1, nil)

		} else {
			if d := arg2.Form.Get("token"); d != "" {
				// 应该判断 获取的这个token的value跟实际我们熟知的该有的token是否相同。
				fmt.Println(d)
			} else {
				plus(arg1, arg2)
				return
			}
			//  如果使用了fromValue的话就不用使用r.ParseForm了。它自动的调用了这个东西。
			if len(arg2.Form["username"][0]) == 0 {
				arg1.Write([]byte(`<html><head><meta charset="utf-8"/></head><body><script>(()=> {alert("请输出信息！")	;})()</script></body></html>`))

			} else {
				t, e := template.ParseFiles("./log-post.html")
				if e != nil {
					fmt.Print(e)
					fmt.Println("在log-post发生错误")
					plus(arg1, arg2)
					return
				}
				fmt.Println(arg2.Form)

				t.Execute(arg1, arg2.Form)
			}
		}

	})
	http.HandleFunc("/uploadfile", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		if r.Method == "GET" {
			unt := time.Now().Unix()
			// md5时间戳才需要使用sum
			m := md5.New()
			io.WriteString(m, strconv.FormatInt(unt, 10))
			token := fmt.Sprintf("%x", m.Sum(nil))
			r1, e := template.ParseFiles("./upload-get.html")
			if e != nil {
				fmt.Println("在get-upload出错")
				plus(w, r)

			}
			r1.Execute(w, token)
		} else {
			r1, e := template.ParseFiles("./uploadfile.html")
			if e != nil {
				fmt.Println("解析uploadfile出错")
				plus(w, r)
			}
			r.ParseMultipartForm(32 << 20)
			file, handler, err := r.FormFile("uploadfile")
			if err != nil {
				fmt.Println("在解析 FormFile时出错")
				plus(w, r)
			}
			defer file.Close()
			// fmt.Fprintf(w, "%x", handler.Header)
			// 这个地方的目的是为了创建一个空的带有名称的名字。
			//    O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
			f, err := os.OpenFile("../../"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err)
				fmt.Println("在openfile处出错")
				plus(w, r)
			}

			defer f.Close()
			io.Copy(f, file)
			r1.Execute(w, handler.Header)
			w.Write([]byte(`<html><head><script>
				(()=>{
					setTimeout(()=>{
						location.href="/"
					},5000)
				})()
				</script></head><body></body></html>`))
		}
	})

	e := http.ListenAndServe(":8080", nil)
	if e != nil {
		fmt.Println(e)
	}

}

func plus(arg1 http.ResponseWriter, arg2 *http.Request) {
	arg1.Write([]byte(`<html><head><script>
		(()=>{
			setTimeout(()=>{
				location.href="/"
			})
		})()
		</script></head><body></body></html>`))
}

func bi(filename, targetURL string) {
	fmt.Println("执行了")
	// 空的缓存
	bodyBuf := &bytes.Buffer{}
	// 新的 multipart writer
	bodyWriter := multipart.NewWriter(bodyBuf)
	// io.writer
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println(err)
	}
	// 打开了文件
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer fh.Close()
	// 拷贝文件到 filewriter中。
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		fmt.Println(err)
	}
	// 获取到content类型
	contentype := bodyWriter.FormDataContentType()
	defer bodyWriter.Close()
	// 使用post传送数据到服务器端的 targeturl处。
	res, e := http.Post(targetURL, contentype, bodyBuf)
	if e != nil {
		fmt.Println(e)
	}
	defer res.Body.Close()
	//
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(resBody))
	fmt.Println(res.Status)

}
