package main

import (
	"context"
	"fmt"
	"time"
)

func reqTask(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name + "done")
			return
		default:
			fmt.Println(name + "wait")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go reqTask(ctx, "task1")
	go reqTask(ctx, "task2")

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)

}
