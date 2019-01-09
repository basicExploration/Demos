//练习 8.8： 使用select来改造8.3节中的echo服务器，为其增加超时，这样服务器可以在客户端10秒中没有任何喊话时自动断开连接。
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
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
		go echo(c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}



func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case <- time.After(time.Second *10):
			return
		default:
			conn, err := l.Accept()
			if err != nil {
				log.Print(err) // e.g., connection aborted
				continue
			}
			go handleConn(conn)
		}

	}
}

//func main() {
//	v := make(chan bool)
//	go tt(v)
//	for {
//		select {
//		case <-time.After(time.Second * 10):
//			return
//		default:
//			fmt.Println(<-v)
//		}
//	}
//
//	time.Sleep(1e10)
//}
//
//func tt(c chan bool) {
//	var t int
//	for {
//		t++
//		if t == 11 {
//			return
//		}
//		c <- true
//		time.Sleep(time.Second*3)
//	}
//
//
//
//}

//func main() {
//	dd()
//	time.Sleep(1e10)
//}
//
//func dd() {
//	dd := make(chan bool)
//	go tttt(dd)
//	ctx, cancle := context.WithDeadline(context.Background(),time.Now().Add(time.Second*3))
//	defer cancle()
//	tt(ctx, dd)
//
//}
//func tttt(d chan  bool){
//	de := 0
//	for {
//		de ++
//		if de == 10 {
//			return
//		}
//		d <- true
//		time.Sleep(time.Second)
//	}
//
//}
//func tt(ctx context.Context, d chan bool) {
//	for {
//		select {
//		case <-ctx.Done():
//			return
//		case <-time.After(time.Second):
//			return
//		default:
//			fmt.Println(<-d)
//		}
//	}
//}
