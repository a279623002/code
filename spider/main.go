package main

import (
	"fmt"
	"sync"
)

type Node struct {
	data int
	next *Node
}

func createCycle() {
	a := &Node{}
	b := &Node{}
	a.next = b
	b.next = a
	// a和b形成循环引用，若后续无其他引用，仍可能导致内存泄漏
}

var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}

func start() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		increment(&wg)
	}
	wg.Wait()
	fmt.Print(x)
}

func t1() int {
    a := 1
    return func () int  {
        return a
    }()
}

func main() {
	// p := platform.NewTaobao("./download/taobao", &sync.WaitGroup{})
	// // p := platform.NewJingdong("./download/jingdong", &sync.WaitGroup{})

	// pm := platform.NewPlatformManager(p)
	// pm.Start()
	// createCycle()
	// start()

}
