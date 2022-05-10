package example

import (
	"my-redis/configs"
	"testing"
)


type testBloomFilter struct {
	b        *BloomFilter
	name     string
	expected bool
}

func createMulTestcase(t *testing.T, tbf *testBloomFilter) {
	if res := tbf.b.Exists(tbf.name); res != tbf.expected {
		t.Fatalf("%s expected %t, but %t got", tbf.name, tbf.expected, res)
	}
}

func TestExists(t *testing.T) {
	session := &configs.SessionConfig{Host:"127.0.0.1:6379",Select:10}
	configs.InitDb(session)

	b := NewBloomFilter("shiro", []string{"PJWHash", "DEKHash", "DJBHash"})
	b.Add("shiro")

	createMulTestcase(t, &testBloomFilter{
		b:        b,
		name:     "shiro",
		expected: true,
	})
	createMulTestcase(t, &testBloomFilter{
		b:        b,
		name:     "hanbaga",
		expected: true,
	})
}
