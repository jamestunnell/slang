package expressions

import "github.com/jamestunnell/slang"

type Equal struct {
	*BinaryOperation
}

func NewEqual(left, right slang.Expression) slang.Expression {
	return &Equal{
		BinaryOperation: NewBinaryOperation(slang.ExprEQUAL, left, right),
	}
}
