package expression

type Expression interface {
	Interpret(context string) bool
}
