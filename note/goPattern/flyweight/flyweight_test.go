package flyweight

import (
	"testing"
)

func TestNew(t *testing.T) {
	if res := New("zzq"); res != "zzq" {
		t.Errorf("res expected be zzq, but %s got", res)
	}
}
