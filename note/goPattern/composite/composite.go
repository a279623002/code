package composite

type Node interface {
	Add(chlid Node)
	Print(space int)
}
