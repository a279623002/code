package factory

type OperatorFactoryPlus struct {
}

type OperatorPlus struct {
	*OperatorBase
}

func (o *OperatorPlus) Result() int {
	return o.Left + o.Right
}

func (op *OperatorFactoryPlus) Create() Operator {
	return &OperatorPlus{&OperatorBase{}}
}
