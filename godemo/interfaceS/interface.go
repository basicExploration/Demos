package main

import "fmt"

type human struct {
	name string
	age  int
}

type stuent struct {
	human
	homework string
}
type teacher struct {
	human
	work string
}

func (a human) sayHi() string {
	// do something
	return "something!"
}
func (a human) sayGood() string {
	// do something
	return "something!"
}
func (b stuent) doHomeWork() string {
	// do something
	return "something!"
}

func (c teacher) doWork() string {
	// do something
	return "something!"
}

type kidWork interface {
	sayGood() string
	sayHi() string
	doHomeWork() string
}
type teacherWork interface {
	sayGood() string
	sayHi() string
	doWork() string
}

func main1() {

	xiaoming := stuent{human{"ad", 12}, "ad"}
	lijuan := teacher{human{"ad", 34}, "ad"}
	var i kidWork
	i = xiaoming
	fmt.Print(i.sayGood(), i.sayHi(), i.doHomeWork())

	var j teacherWork
	j = lijuan
	fmt.Print(
		j.sayGood(), j.sayHi(), j.doWork())

}

// 所以说interface就是很多方法的集合。它的主要的作用就是这个。要记得。
