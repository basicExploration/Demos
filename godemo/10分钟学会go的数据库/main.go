// Package main 这个项目是为了操作go的数据库，mgo。

package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone int
}

func main() {
	session, err := mgo.Dial("mongodb://userName:password@hostname:port/dbname")
	if err != nil {
		fmt.Println("连接错误")
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("cr").C("u")
	for i := 0; i < 10; i++ {
		err = c.Insert(&Person{"Thomashuke", i},
			&Person{"mylove", i + 10})
		if err != nil {
			fmt.Println("错误了，错误的信息是", err, "错误的地址是插入的时候")
			log.Fatal(err)
		}

	}

	result := Person{}
	err = c.Find(bson.M{"name": "Thomashuke"}).One(&result)
	if err != nil {
		fmt.Println("错误了，错误的信息是", err, "错误的地址是查询的时候的时候")
		log.Fatal(err)
	}

	fmt.Println("姓名是:", result.Phone)
}
