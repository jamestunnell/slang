package expressions

import "github.com/jamestunnell/slang"

type Subtract struct {
	*BinaryOperation
}

func NewSubtract(left, right slang.Expression) slang.Expression {
	return &Subtract{
		BinaryOperation: NewBinaryOperation(slang.ExprSUBTRACT, left, right),
	}
}
