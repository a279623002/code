package factory

type OperatorBase struct {
	Left, Right int
}

func (o *OperatorBase) SetLeft(left int) {
	o.Left = left
}

func (o *OperatorBase) SetRight(right int) {
	o.Right = right
}
