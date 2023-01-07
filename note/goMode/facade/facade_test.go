package facade

import (
	"testing"
)

func TestNewShapeMaker(t *testing.T) {
	s := NewShapeMaker()
	if res := s.circle.Draw(); res != "Circle" {
		t.Errorf("res expected be Circle, but %s got", res)
	}
	if res := s.rectangle.Draw(); res != "Rectangle" {
		t.Errorf("res expected be Rectangle, but %s got", res)
	}
	if res := s.square.Draw(); res != "Square" {
		t.Errorf("res expected be Square, but %s got", res)
	}
}
