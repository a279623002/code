package main

import "interfaceModel/shiro"

func main() {
	s := shiro.New()
	s.SetAdapter(shiro.NewAdapterZzq("hello, zzq"))
	s.Hello()
}
