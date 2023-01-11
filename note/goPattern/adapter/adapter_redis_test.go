package adapter

import "testing"

func TestAdapterRedis_Get(t *testing.T) {
	cache := NewCache()

	redis := NewAdapterRedis()
	cache.SetAdapter(redis)

	cache.Set("shiro")

	if res := cache.Get(); res != "shiro" {
		t.Errorf("res expected be shiro, but %s got", res)
	}
}
