package prototype

import (
	"testing"
)

func TestNewPrototypeManage(t *testing.T) {
	m := NewPrototypeManage()
	s := &Shiro{Name: "shiro"}
	m.ClientArr["shiro"] = s

	ss := m.ClientArr["shiro"].Clone()
	if s == ss {
		t.Errorf("res be s != ss, but s == ss got")
	}
}
