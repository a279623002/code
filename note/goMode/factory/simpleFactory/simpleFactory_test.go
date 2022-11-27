package simpleFactory

import (
	"testing"
)

func TestChina_Say(t *testing.T) {
	cn := NewDrive("cn")

	if res := cn.Say("1"); res != "cn:1" {
		t.Errorf("res expected be cn:1, but %s got", res)
	}
}

func TestEnglish_Say(t *testing.T) {
	cn := NewDrive("en")

	if res := cn.Say("1"); res != "en:1" {
		t.Errorf("res expected be en:1, but %s got", res)
	}
}
