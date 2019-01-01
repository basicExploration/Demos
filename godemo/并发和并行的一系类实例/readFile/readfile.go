// io 型 bound
package main

import (
	"encoding/json"
	"sync"
	"time"
)

func read(s string) {
	document := new(jsonValue)
	time.Sleep(time.Millisecond)
	json.Unmarshal([]byte(s), document)
}

func find(doc []string) {
	for _, v := range doc {
		read(v)
	}
}

func findCon(cpu int, doc []string) {
	var synv sync.WaitGroup
	ch := make(chan string, len(doc))
	for i := 0; i < len(doc); i++ {
		ch <- doc[i]
	}
	close(ch)
	synv.Add(cpu)
	for i := 0; i < cpu; i++ {
		go func() {
			for v := range ch {
				read(v)
			}
			synv.Done()
		}()
	}

	synv.Wait()

}
