package main

import (
	"context"
	"fmt"
	"time"
)

func main(){
ctx,cancel := context.WithDeadline(context.Background(),time.Now().Add(3*time.Second))
ctx1 := context.WithValue(ctx,"12","[\"监控开始1\"]")// 对下属的context进行传递信息
ctx2 := context.WithValue(ctx,"12","[二代监控]")
go read(ctx1)
go read1(ctx2)
time.Sleep(10*time.Second)
cancel()
}

func read(ctx context.Context){
	for   {
		time.Sleep(1*time.Second)
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value("12"))
			fmt.Println(time.Now().Nanosecond())
			fmt.Println("done")
		default:
			fmt.Println(time.Now().Nanosecond())
			fmt.Println("working!")

		}
	}
}

func read1(ctx context.Context){
	for {
		time.Sleep(time.Second)
		select {
		case <- ctx.Done():
			fmt.Println(ctx.Value("12"))
			fmt.Println("done")
		default:
			fmt.Println("work")
		}
	}
}
