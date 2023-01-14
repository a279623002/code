package observer

import "fmt"

// 观察者
type aObserver struct {
	name string
}

func (o *aObserver) Receive(str string) {
	fmt.Println("A观察者[" + o.name + "]接受:" + str)
}

type bObserver struct {
	name string
}

func (o *bObserver) Receive(str string) {
	fmt.Println("B观察者[" + o.name + "]接受:" + str)
}
