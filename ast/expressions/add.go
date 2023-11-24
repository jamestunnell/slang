package expressions

import "github.com/jamestunnell/slang"

type Add struct {
	*BinaryOperation
}

func NewAdd(left, right slang.Expression) slang.Expression {
	return &Add{
		BinaryOperation: NewBinaryOperation(slang.ExprADD, left, right),
	}
}
