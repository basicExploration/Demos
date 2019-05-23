// 链表队列
package main

import (
	"fmt"
	"sync"
)

// 使用单链表就ok了。
type node struct {
	Value interface{}
	Next  *node
}

type linkedListQueue struct {
	lock   sync.Mutex
	Length int
	head   *node //头节点
	tail   *node //尾节点
}

func NewLinkListQueue() *linkedListQueue {
	return &linkedListQueue{
		Length: 0,
		head:   nil,
		tail:   nil,
	}
}

func (l *linkedListQueue) Traverse() []interface{} {
	b := l.head
	result := make([]interface{}, 0)
	for {
		result = append(result, b.Value)
		b = b.Next
		if b.Next == nil {
			result = append(result,b.Value)
			break
		}
	}
	return result
}
func (l *linkedListQueue) Enqueue(v interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	newNode := &node{
		Value: v,
		Next:  nil,
	}
	if l.Length == 0 {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.Next = newNode // 将本次的tail的next指向newNode
		l.tail = newNode      // 将tail换人换成 这个newnode
	}
	l.Length++
}

func (l *linkedListQueue) Dequeue() (interface{}, error) {
	l.lock.Lock()
	defer l.lock.Unlock()
	var item interface{}
	if l.Length == 0 {
		return "", fmt.Errorf("空")
	} else if l.Length == 1 {
		return l.head.Value, nil
	} else {
		item = l.head.Value
		l.head = l.head.Next
	}
	l.Length--
	return item, nil
}

func main() {
	lq := NewLinkListQueue()
	lq.Enqueue(12)
	lq.Enqueue(13)
	lq.Enqueue(14)
	lq.Enqueue(15)
	lq.Enqueue(16)
	lq.Enqueue(17)
	lq.Dequeue()
	fmt.Println(lq.Traverse())

}
