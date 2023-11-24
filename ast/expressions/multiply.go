package expressions

import "github.com/jamestunnell/slang"

type Multiply struct {
	*BinaryOperation
}

func NewMultiply(left, right slang.Expression) slang.Expression {
	return &Multiply{
		BinaryOperation: NewBinaryOperation(slang.ExprMULTIPLY, left, right),
	}
}
