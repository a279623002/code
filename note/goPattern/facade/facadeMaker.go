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

func (s *ShapeMaker) DrawCircle() string {
	return s.circle.Draw()
}

func (s *ShapeMaker) DrawSquare() string {
	return s.square.Draw()
}

func (s *ShapeMaker) DrawRectangle() string {
	return s.rectangle.Draw()
}
