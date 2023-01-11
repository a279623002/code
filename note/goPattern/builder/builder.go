package builder

// 按步骤构造，如启动一个服务，先传入配置初始->连接网络->启动服务
type Builder interface {
	Step1()
	Step2()
	Step3()
}

type Director struct {
	builder Builder
}

func NewDirector(b Builder) *Director {
	return &Director{b}
}

func (d *Director) MakeData() {
	d.builder.Step1()
	d.builder.Step2()
	d.builder.Step3()
}
