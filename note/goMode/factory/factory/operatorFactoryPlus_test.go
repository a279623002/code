package factory

import "testing"

func TestOperatorPlus_Result(t *testing.T) {
	var fac OperatorFactory
	fac = &OperatorFactoryPlus{}
	op := fac.Create()
	op.SetLeft(1)
	op.SetRight(2)

	if res := op.Result(); res != 3 {
		t.Errorf("res expected be 3, but %d got", res)
	}
}
