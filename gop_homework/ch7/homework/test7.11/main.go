//练习 7.11： 增加额外的handler让客服端可以创建，读取，更新和删除数据库记录。
// 例如，一个形如 /update?name=1&num=1 的请求会更新一个人的name和num
// 并且当这个人的名字无效时返回一个错误值。
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
	"strings"
)

var (
	db  *sql.DB
	err error
)

func init() {
	db, err = sql.Open("mysql", "#")
	if err != nil {
		fmt.Println("连接数据库发生错误：", err)
		os.Exit(1)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.ToUpper(r.URL.Path) == "/" {
			var sqldata string
			row, err := db.Query("SELECT name FROM room")
			if err != nil {
				fmt.Println(err)
			}
			for row.Next() {
				t := new(string)
				err = row.Scan(t)
				if err != nil {
					fmt.Println(err)
				}
				sqldata += *t
			}
			fmt.Println("d")
			fmt.Println("ceshi", sqldata)
			w.Write([]byte(sqldata))
		}
	})
	http.HandleFunc("/up", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		number := r.FormValue("number")
		db.QueryRow("UPDATE room SET name=?", name).Scan(nil)
		db.QueryRow("UPDATE room SET number=?", number).Scan(nil)
	})
	http.ListenAndServe(":8090", nil)
}
