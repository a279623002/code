package main

import (
	"fmt"
	"time"
)

var (
	ch1 = make(chan string)
	ch2 = make(chan string)
	exit = make(chan bool)
)

func push(ch chan string, msg string) {
	ch <- msg
}

func do() {
	for {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		case <-exit:
			fmt.Println("exit")
			return
		default:
			fmt.Println("no msg")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			push(ch1, fmt.Sprintf("ch1 msg:%d", i))
			push(ch2, fmt.Sprintf("ch2 msg:%d", i))
		}
		exit<-true
	}()

	do()
	fmt.Println("done")
}
