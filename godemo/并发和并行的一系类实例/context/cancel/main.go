package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx,cancel := context.WithCancel(context.Background())
	go read(ctx)
	go read1(ctx)
	time.Sleep(13*time.Second)
	cancel()
}

func read(ctx context.Context) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("done")
		default:
			i++
			fmt.Println("work")
		}
		fmt.Println("i",i)
	}
}

func read1(ctx context.Context){
	i := 0
	for {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("done")
		default:
			fmt.Println("work1")
			i++
		}
		fmt.Println("i-1",i)
	}

}