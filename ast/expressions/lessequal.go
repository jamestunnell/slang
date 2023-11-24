package expressions

import "github.com/jamestunnell/slang"

type LessEqual struct {
	*BinaryOperation
}

func NewLessEqual(left, right slang.Expression) slang.Expression {
	return &LessEqual{
		BinaryOperation: NewBinaryOperation(slang.ExprLESSEQUAL, left, right),
	}
}
