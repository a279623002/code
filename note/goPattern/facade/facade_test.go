package facade

import (
	"testing"
)

func TestNewShapeMaker(t *testing.T) {
	s := NewShapeMaker()
	if res := s.DrawCircle(); res != "Circle" {
		t.Errorf("res expected be Circle, but %s got", res)
	}
	if res := s.DrawRectangle(); res != "Rectangle" {
		t.Errorf("res expected be Rectangle, but %s got", res)
	}
	if res := s.DrawSquare(); res != "Square" {
		t.Errorf("res expected be Square, but %s got", res)
	}
}
