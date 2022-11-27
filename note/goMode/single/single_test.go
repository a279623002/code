package single

import "testing"

func TestGetHandler(t *testing.T) {
	s1 := GetHandler()
	s2 := GetHandler()

	if s1 != s2 {
		t.Errorf("res expected be s1 == s2, but s1 != s2 got")
	}
}
