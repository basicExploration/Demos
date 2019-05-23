package main

import (
	"fmt"
	"sync"
)

// 循环队列
type LoopQueue struct {
	lock  sync.Mutex
	data  []interface{}
	front int
	trail int
	l     int
	cap   int
}

func (l *LoopQueue) len() int {
	return l.l
}
func NewLoopQueue(cap int) *LoopQueue {
	l := &LoopQueue{
		data:  make([]interface{}, 0, cap),
		front: 0,
		trail: 0,
		l:     0,
		cap:   cap,
	}
	for i := 0; i < cap; i++ {
		l.data = append(l.data, "")
	}
	return l
}

// 入
func (l *LoopQueue) EnQueue(t interface{}) error {
	l.lock.Lock()
	defer l.lock.Unlock()
	if (l.trail+1)%l.cap == (l.front)%l.cap {
		return fmt.Errorf("满了")
	}
	l.data[l.trail] = t
	l.trail = (l.trail + 1) % l.cap
	l.l++
	return nil
}

// 出
func (l *LoopQueue) DeQueue() error {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.front == l.trail {
		return fmt.Errorf("没有东西")
	}
	l.data[l.front] = 0
	l.front = (l.front + 1) % l.cap
	return nil
}

func main() {
	l := NewLoopQueue(10)
	l.EnQueue(1)
	l.EnQueue(1)
	fmt.Println(l.len())
	l.EnQueue(1)
	l.EnQueue(1)
	l.DeQueue()
	l.EnQueue(1)
	l.EnQueue(1)
	l.DeQueue()
	l.DeQueue()
	l.DeQueue()
	l.EnQueue(1)
	l.EnQueue(1)
	l.EnQueue(1)
	l.EnQueue(1)
	l.EnQueue(1)
	l.EnQueue(1)
	l.EnQueue(1)
	l.EnQueue(1)
	fmt.Println(l.data)
}
