package main

import (
	"fmt"
	"sync"
)

type AbsentStruct struct {
	Name string
}

var Absent *AbsentStruct
var once = sync.Once{}

func AbsentInstance() *AbsentStruct {
	once.Do(func() {
		Absent = new(AbsentStruct)
	})
	return Absent
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			a1 := AbsentInstance()
			fmt.Printf("%p \n %s \n", a1, a1.Name)
		}()
	}
	wg.Wait()
}
