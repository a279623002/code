package factory

type Operator interface {
	SetLeft(int)
	SetRight(int)
	Result() int
}

type OperatorFactory interface {
	Create() Operator
}
