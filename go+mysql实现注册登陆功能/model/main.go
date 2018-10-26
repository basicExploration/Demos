package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 储存在数据库中的的用户信息
type Result struct {
	Username string
	Password string
	B        BoolResult
}

// 判断结果
type BoolResult struct {
	B bool
}

// 开启数据库
func NewDb() *sql.DB {
	fmt.Println("NewDb 正在执行")
	db, err := sql.Open("mysql", "root:359258lS@/myuser?charset=utf8")
	if err != nil {
		Where(err)
	}
	return db
}

// 添加数据到数据库
func (r Result) Add() {

	db := NewDb()
	fmt.Println("正常开启数据库")
	defer db.Close()
	stmt, err := db.Prepare("INSERT userinfo set username=?,password=?")
	if err != nil {
		Where(err)
	}

	a, b := r.Username, r.Password
	r.B.B = false
	fmt.Println("add中添加的数据是：", a, b)
	res, err := stmt.Exec(a, b)
	WherePlus(err)
	id, err := res.LastInsertId()
	if err != nil {
		Where(err)
	}
	fmt.Println("已经添加数据库中数据是：", id)
	return

}

// 从数据库中查找数据
func (r Result) Find(username, password string) bool {
	fmt.Println("查找数据库正在执行")
	db := NewDb()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM userinfo")
	WherePlus(err)
	for rows.Next() {
		var uid int
		var userna string
		var passwo string
		err = rows.Scan(&uid, &userna, &passwo)
		fmt.Println("find中查询的数据是:", userna, passwo)
		if err != nil {
			Where(err)
		}
		if username == userna {
			if passwo == password {
				return true
			} else {
				continue
			}
		}
	}
	return false
}
