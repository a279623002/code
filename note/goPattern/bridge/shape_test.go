package bridge

import "testing"

func TestNewShapeCircle(t *testing.T) {

	s := NewShapeCircle(5, NewRedCircle())
	if s.drawAPI.DrawCircle(s.Radius) != 5 {
		t.Errorf("res be s == 5, but s != ss got")
	}
}
