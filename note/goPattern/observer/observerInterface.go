package observer

type Subject interface {
	Add(o Observer)
	Send(str string)
}

type Observer interface {
	Receive(str string)
}
