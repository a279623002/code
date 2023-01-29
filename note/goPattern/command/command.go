package command

// 抽象的命令
type Command interface {
	Treat()
}

// 治疗眼睛
type CommandTreatEye struct {
	d *doctor
}

func (c *CommandTreatEye) Treat() {
	c.d.treatEye()
}

// 治疗鼻子
type CommandTreatNose struct {
	d *doctor
}

func (c *CommandTreatNose) Treat() {
	c.d.treatNose()
}

// 护士-调用命令者
type Nurse struct {
	CmdList []Command
}

func (n *Nurse) Notify() {
	if n.CmdList == nil {
		return
	}
	for _, cmd := range n.CmdList {
		cmd.Treat()
	}
}

func Run() {
	d := &doctor{}
	cmdEye := &CommandTreatEye{d: d}
	cmdNose := &CommandTreatNose{d: d}
	n := &Nurse{}
	n.CmdList = append(n.CmdList, cmdEye)
	n.CmdList = append(n.CmdList, cmdNose)
	n.Notify()
}
