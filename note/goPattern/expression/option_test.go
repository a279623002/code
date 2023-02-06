package expression

import (
	"fmt"
	"testing"
)

func TestOrExpression_Interpret(t *testing.T) {
	//规则，lee 和 wang 是男性
	lee := NewTerminalExpression("Lee")
	wang := NewTerminalExpression("Wang")
	isMale := NewOrExpression(lee, wang)

	fmt.Println("lee is male? ", isMale.Interpret("Lee"))
}

func TestAndExpression_Interpret(t *testing.T) {
	//规则，yang是已婚女性
	yang := NewTerminalExpression("Yang")
	married := NewTerminalExpression("Married")
	isMarried := NewAndExpression(yang, married)

	fmt.Println("Yang is a married women? ", isMarried.Interpret("Married Yang"))
}
