package expression

import "strings"

type TerminalExpression struct {
	Data string
}

func NewTerminalExpression(data string) *TerminalExpression {
	return &TerminalExpression{Data: data}
}

func (t *TerminalExpression) Interpret(context string) bool {
	if strings.Contains(context, t.Data) {
		return true
	}
	return false
}
