package builder

type BuilderStr struct {
	res string
}

func (bs *BuilderStr) Step1() {
	bs.res += "a"
}

func (bs *BuilderStr) Step2() {
	bs.res += "b"
}

func (bs *BuilderStr) Step3() {
	bs.res += "c"
}

func (bs *BuilderStr) Res() string {
	return bs.res
}
