package main

import (
	"fmt"
	"time"
)

var ch = make(chan string)

func goroutine(i int) {
	msg := fmt.Sprintf("goroutine run %d\n", i)
	fmt.Println(msg)
	time.Sleep(time.Second)
	ch <- msg
}

func main() {
	for i := 0; i < 4; i++ {
		go goroutine(i)
	}

	for i := 0; i < 4; i++ {
		msg := <-ch
		fmt.Println("done " + msg)
	}
	fmt.Println("done")
}
