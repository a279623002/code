package bridge

//ShapeCircle 桥接模式的实体类
type ShapeCircle struct {
	Radius  int
	drawAPI DrawApi
}

//NewShapeCircle 实例化桥接模式实体类
func NewShapeCircle(radius int, drawAPI DrawApi) *ShapeCircle {
	return &ShapeCircle{
		Radius:  radius,
		drawAPI: drawAPI,
	}
}

//Draw 实体类的Draw方法
func (sc *ShapeCircle) Draw() int {
	return sc.drawAPI.DrawCircle(sc.Radius)
}
