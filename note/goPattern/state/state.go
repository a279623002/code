package state

// 抽象的状态 功能接口
type State interface {
	ShowState()
	NextState(ctx *context)
}

// 环境对象
type context struct {
	state State
}

func NewContext(initState int) context {
	return context{state: &State1{state: initState}}
}

func (c *context) Start() {
	c.state.ShowState()
	c.state.NextState(c)
}

func Run() {
	ctx := NewContext(1)
	for i := 0; i <= 5; i++ {
		ctx.Start()
	}
}
