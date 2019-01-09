//练习 8.13： 使聊天服务器能够断开空闲的客户端连接，比如最近五分钟之后没有发送任何消息的那些客户端
// 。提示：可以在其它goroutine中调用conn.Close()来解除Read调用，就像input.Scanner()所做的那样。
package main

import "time"

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case <-time.After(time.Minute * 5):
			return // 已经延迟5分钟了，可以关闭这个连接了。
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}