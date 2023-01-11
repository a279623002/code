package decorator

type Beverage interface {
	getDescription() string
	cost() int
}

type Coffee struct {
	desc string
}

func (c *Coffee) getDescription() string {
	return c.desc
}

func (c *Coffee) cost() int {
	return 1
}

type MoCcha struct {
	bev  Beverage
	desc string
}

func (m *MoCcha) getDescription() string {
	return m.bev.getDescription() + m.desc
}

func (m *MoCcha) cost() int {
	return m.bev.cost() + 1
}

type Whip struct {
	bev  Beverage
	desc string
}

func (w *Whip) getDescription() string {
	return w.bev.getDescription() + w.desc
}

func (w *Whip) cost() int {
	return w.bev.cost() + 1
}

func Run() (desc string, cost int) {
	var bev Beverage
	bev = &Coffee{desc: "c"}
	bev = &MoCcha{bev: bev, desc: "m"}
	bev = &Whip{bev: bev, desc: "w"}
	return bev.getDescription(), bev.cost()
}
