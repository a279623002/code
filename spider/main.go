package main

func main() {
	// p := platform.NewTaobao("./download/taobao", &sync.WaitGroup{})
	// // p := platform.NewJingdong("./download/jingdong", &sync.WaitGroup{})

	// pm := platform.NewPlatformManager(p)
	// pm.Start()
	s := []int{1, 2, 3, 4, 5}
	for k, _ := range s {
		s[k] = 0
	}
	for _, v := range s {
		print(v)
	}
}
