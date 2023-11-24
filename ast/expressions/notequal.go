package expressions

import "github.com/jamestunnell/slang"

type NotEqual struct {
	*BinaryOperation
}

func NewNotEqual(left, right slang.Expression) slang.Expression {
	return &NotEqual{
		BinaryOperation: NewBinaryOperation(slang.ExprNOTEQUAL, left, right),
	}
}
