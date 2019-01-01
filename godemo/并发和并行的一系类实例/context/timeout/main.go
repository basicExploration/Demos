package main

import (
	"fmt"
	"time"
	"context"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	go read(ctx)
	time.Sleep(10 * time.Second)
	cancel()
}

func read(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println(time.Now().Nanosecond())
			fmt.Println("done")
		default:
			fmt.Println(time.Now().Nanosecond())
			fmt.Println("working!")

		}
	}
}
