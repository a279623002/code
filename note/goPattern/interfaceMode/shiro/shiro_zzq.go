package shiro

import "fmt"

type AdapterZzq struct {
	Name string
}

func NewAdapterZzq(name string) *AdapterZzq {
	return &AdapterZzq{Name: name}
}

func (a *AdapterZzq) Hello() {
	fmt.Println(a.Name)
}
