package facade

//模型实例创建
type Shape interface {
	Draw() string
}

type Circle struct {
}

type Rectangle struct {
}

type Square struct {
}

func NewCircle() *Circle {
	return &Circle{}
}

func (c *Circle) Draw() string {
	return "Circle"
}

func NewRectangle() *Rectangle {
	return &Rectangle{}
}

func (r *Rectangle) Draw() string {
	return "Rectangle"
}

func NewSquare() *Square {
	return &Square{}
}

func (s *Square) Draw() string {
	return "Square"
}
