package main

import "sync"

func main() {
	title(map[string]string{
		"1": "1",
	})

}

var lock sync.Mutex

func title(k map[string]string) map[string]string {
	lock.Lock()
	defer lock.Unlock()
	return k
}

// 说明几点：1，互斥锁可以在不同的线程中打开和关闭，例如在a线程中打开在b线程中关闭，这是可以的。2，普通的map在多线程中是有访问冲突的，需要使用互斥锁，
//或者你可以使用sync.map这个是sync包实现的map估计实现原理就跟一般的map加锁差不多。3有关互斥锁，file html访问这种情境下一般而言都要使用defer，这样
//结构更清晰也更好用。
