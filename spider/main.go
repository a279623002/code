package main


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


func main() {
	// p := platform.NewTaobao("./download/taobao", &sync.WaitGroup{})
	// // p := platform.NewJingdong("./download/jingdong", &sync.WaitGroup{})

	// pm := platform.NewPlatformManager(p)
	// pm.Start()
	createCycle()
}
