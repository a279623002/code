package proxy

import "testing"

func TestProxy_Do(t *testing.T) {
	proxy := &Proxy{
		rel:   &RelSubject{},
		money: 100,
	}
	if res := proxy.Do(); res != "hello" {
		t.Errorf("res expected be hello, but %s got", res)
	}
}
