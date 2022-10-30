package utils

import "fmt"

type Shiro interface {
	Name() string
	Set(string)
}

func Show(shiro Shiro) {
	fmt.Println(shiro.Name())
}
