package main

import (
	"fmt"
	"sync"
)

func main() {
	t := NewThing()
	t.In("12")
	t.In("13")
	t.In("14")
	t.In("15")
	t.Out()
	fmt.Println(t.value)
}

type thing struct {
	lock sync.Mutex
	value []interface{}
}

func NewThing() *thing {
	return &thing{
		value:make([]interface{}, 0),
	}
}

func (t *thing) In(v interface{}) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.value = append(t.value, v)
}

func (t *thing) Out() interface{} {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.value = t.value[1:]
	return t.value[0]
}



