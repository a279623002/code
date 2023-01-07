package facade

// 外观类实现
type ShapeMaker struct {
	circle    *Circle
	square    *Square
	rectangle *Rectangle
}

func NewShapeMaker() *ShapeMaker {
	return &ShapeMaker{
		circle:    NewCircle(),
		square:    NewSquare(),
		rectangle: NewRectangle(),
	}
}

func (s *ShapeMaker) DrawCircle() {
	s.circle.Draw()
}

func (s *ShapeMaker) DrawSquare() {
	s.square.Draw()
}

func (s *ShapeMaker) DrawRectangle() {
	s.rectangle.Draw()
}
