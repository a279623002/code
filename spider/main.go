package main

import (
	"spider/platform"
	"sync"
)

func main() {
	p := platform.NewTaobao("./download/taobao", &sync.WaitGroup{})
	// p := platform.NewJingdong("./download/jingdong", &sync.WaitGroup{})

	pm := platform.NewPlatformManager(p)
	pm.Start()
}
