package state

import "fmt"

// 具体的 状态类
type State1 struct {
	state int
}

func (c *State1) ShowState() {
	fmt.Println(c.state)
}

func (c *State1) NextState(ctx *context) {
	ctx.state = &State2{state: 2}
}

type State2 struct {
	state int
}

func (c *State2) ShowState() {
	fmt.Println(c.state)
}

func (c *State2) NextState(ctx *context) {
	ctx.state = &State3{state: 3}
}

type State3 struct {
	state int
}

func (c *State3) ShowState() {
	fmt.Println(c.state)
}

func (c *State3) NextState(ctx *context) {
	ctx.state = &State1{state: 1}
}
