package builder

type BuilderInt struct {
	res int
}

func (bs *BuilderInt) Step1() {
	bs.res += 1
}

func (bs *BuilderInt) Step2() {
	bs.res += 2
}

func (bs *BuilderInt) Step3() {
	bs.res += 3
}

func (bs *BuilderInt) Res() int {
	return bs.res
}
