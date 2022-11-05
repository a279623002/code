package factory

type OperatorFactorySub struct {
}

type OperatorSub struct {
	*OperatorBase
}

func (o *OperatorSub) Result() int {
	return o.Left - o.Right
}

func (op *OperatorFactorySub) Create() Operator {
	return &OperatorSub{&OperatorBase{}}
}
