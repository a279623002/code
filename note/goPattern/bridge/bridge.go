package bridge

//DrawAPI 画图抽象接口，桥接模式的抽象接口
type DrawApi interface {
	DrawCircle(radius int) int
}

//RedCircle 红色圆的类，桥接模式接口
type RedCircle struct{}

//NewRedCircle 实例化红色圆
func NewRedCircle() *RedCircle {
	return &RedCircle{}
}

//DrawCircle 红色圆实现DrawAPI方法
func (rc *RedCircle) DrawCircle(radius int) int {
	return radius
}
