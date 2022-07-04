package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan struct{}, 3)
	for i:=0; i<10000; i++ {
		wg.Add(1)
		ch<-struct{}{}
		go func(i int) {
			log.Println(i)
			time.Sleep(time.Second)
			<-ch
			wg.Done()
		}(i)
	}
	wg.Wait()
}
