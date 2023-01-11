package flyweight

// 共享的享元类型
type Flyweight struct {
	Data string
}

// 资源工厂
type FlyweightFactory struct {
	memory map[string]*Flyweight
}

func (f *FlyweightFactory) GetFlyweight(name string) *Flyweight {
	flyweight, ok := f.memory[name]
	if !ok {
		flyweight = &Flyweight{Data: name}
		f.memory[name] = flyweight
	}
	return flyweight
}

var factory *FlyweightFactory

func GetFactory() *FlyweightFactory {
	if factory == nil {
		factory = &FlyweightFactory{memory: make(map[string]*Flyweight)}
	}
	return factory
}

// 具体的享元类型
func New(name string) string {
	flyweightFactory := GetFactory()
	flyweight := flyweightFactory.GetFlyweight(name)
	return flyweight.Data
}
