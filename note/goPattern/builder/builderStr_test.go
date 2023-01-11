package builder

import "testing"

func TestBuilderStr_Res(t *testing.T) {
	bs := &BuilderStr{}
	d := NewDirector(bs)
	d.MakeData()
	if res := bs.Res(); res != "abc" {
		t.Errorf("res expected be abc, but %s got", res)
	}
}
