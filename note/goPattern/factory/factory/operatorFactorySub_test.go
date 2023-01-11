package factory

import "testing"

func TestOperatorSub_Result(t *testing.T) {
	var fac OperatorFactory
	fac = &OperatorFactorySub{}
	op := fac.Create()
	op.SetLeft(3)
	op.SetRight(2)

	if res := op.Result(); res != 1 {
		t.Errorf("res expected be 1, but %d got", res)
	}
}
