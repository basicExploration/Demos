//练习 8.4： 修改reverb2服务器，在每一个连接中使用sync.WaitGroup
// 来计数活跃的echo goroutine。
// 当计数减为零时，关闭TCP连接的写入，像练习8.3中一样。
// 验证一下你的修改版netcat3客户端会一直等待所有的并发“喊叫”完成，
// 即使是在标准输入流已经关闭的情况下。
package main
