package utils

import "fmt"

type Zzq struct {
	name string
	age  int
}

func (z *Zzq) Set(str string) {
	z.name = str
}

func (z *Zzq) Name() string {
	return z.name
}

func Print() {
	zzq := &Zzq{}
	zzq.Set("hello, zzq")
	Show(zzq)
	fmt.Println(22333)
}
