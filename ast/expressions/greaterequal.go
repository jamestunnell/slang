package expressions

import "github.com/jamestunnell/slang"

type GreaterEqual struct {
	*BinaryOperation
}

func NewGreaterEqual(left, right slang.Expression) slang.Expression {
	return &GreaterEqual{
		BinaryOperation: NewBinaryOperation(slang.ExprGREATEREQUAL, left, right),
	}
}
