package facade

import (
	"testing"
)

func TestAPICall_Test(t *testing.T) {
	api := &APICall{}
	api.a = NewAModuleApi()
	api.b = NewBModuleApi()
	if res := api.Test(); res != "shiro and zzq" {
		t.Errorf("res expected be shiro and zzq, but %s got", res)
	}
}
