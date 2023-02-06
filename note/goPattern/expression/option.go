package expression

type OrExpression struct {
	Expr1 Expression
	Expr2 Expression
}

func NewOrExpression(expr1, expr2 Expression) *OrExpression {
	return &OrExpression{
		Expr1: expr1,
		Expr2: expr2,
	}
}

func (o *OrExpression) Interpret(context string) bool {
	return o.Expr1.Interpret(context) || o.Expr2.Interpret(context)
}

type AndExpression struct {
	Expr1 Expression
	Expr2 Expression
}

func NewAndExpression(expr1, expr2 Expression) *AndExpression {
	return &AndExpression{
		Expr1: expr1,
		Expr2: expr2,
	}
}

func (o *AndExpression) Interpret(context string) bool {
	return o.Expr1.Interpret(context) && o.Expr2.Interpret(context)
}
