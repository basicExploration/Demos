//练习 8.4： 修改reverb2服务器，在每一个连接中使用sync.WaitGroup
// 来计数活跃的echo goroutine。
// 当计数减为零时，关闭TCP连接的写入，像练习8.3中一样。
// 验证一下你的修改版netcat3客户端会一直等待所有的并发“喊叫”完成，
// 即使是在标准输入流已经关闭的情况下。
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)
var sy sync.WaitGroup
func echo(c net.Conn, shout string, delay time.Duration) {
	defer sy.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		sy.Add(1)
		go echo(c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

//!-

func main() {

	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
	sy.Wait()
}


