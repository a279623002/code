package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := sync.RWMutex{}
	a := 1
	ch := make(chan struct{})

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 5)
			m.Lock()
			a++
			m.Unlock()
		}
		ch <- struct{}{}
	}()

	go func() {
		for {
			m.RLock()
			fmt.Println(a)
			m.RUnlock()
		}
	}()

	<-ch
	fmt.Println("done")

}
