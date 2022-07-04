package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func goroutine1() {
	for i := 0; i < 5; i++ {
		fmt.Printf("goroutine1 run %d\n", i)
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

func goroutine2() {
	for i := 0; i < 5; i++ {
		fmt.Printf("goroutine2 run %d\n", i)
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go goroutine1()
	go goroutine2()

	wg.Wait()
	fmt.Println("done")

}
