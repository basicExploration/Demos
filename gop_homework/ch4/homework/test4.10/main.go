//练习 4.10： 修改issues程序，根据问题的时间进行分类，比如不到一个月的、不到一年的、超过一年。
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

//!+
type Value struct {
	number int
	login  string
	title  string
	time   time.Time
}

func main() {
	t1 := make([]*Value, 0)
	t2 := make([]*Value, 0)
	t3 := make([]*Value, 0)
	now := time.Now()
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		data := now.Sub(item.CreatedAt).Hours() / (24 * 30)
		if data <= 1 {
			v := &Value{
				item.Number,
				item.User.Login,
				item.Title,
				item.CreatedAt,
			}
			t1 = append(t1, v)
		} else if data <= 12 {
			v := &Value{
				item.Number,
				item.User.Login,
				item.Title,
				item.CreatedAt,
			}
			t2 = append(t2, v)
		} else if data > 12 {
			v := &Value{
				item.Number,
				item.User.Login,
				item.Title,
				item.CreatedAt,
			}
			t3 = append(t3, v)
		}
		//fmt.Printf("#%-5d %9.9s %.55s,%s\n",
		//	item.Number, item.User.Login, item.Title, item.CreatedAt.String())
	}
	for _, v := range t1 {// 这所以用指针 这里就是证明如果不用的话 v就会赋值一遍然后被取值。
		fmt.Println("时间：", v.time.String(), "数量: ", v.number, "login: ", v.login, "title: ", v.title)
	}

	fmt.Println("--------------- 这是分界线 --------------")
	for _, v := range t2 {
		fmt.Println("时间：", v.time, "数量: ", v.number, "login: ", v.login, "title: ", v.title)
	}
	fmt.Println("--------------- 这是分界线 --------------")
	for _, v := range t3 {
		fmt.Println("时间：", v.time, "数量: ", v.number, "login: ", v.login, "title: ", v.title)
	}
}
