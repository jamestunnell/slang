package expressions

import "github.com/jamestunnell/slang"

type Divide struct {
	*BinaryOperation
}

func NewDivide(left, right slang.Expression) slang.Expression {
	return &Divide{
		BinaryOperation: NewBinaryOperation(slang.ExprDIVIDE, left, right),
	}
}
