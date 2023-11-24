package expressions

import "github.com/jamestunnell/slang"

type Less struct {
	*BinaryOperation
}

func NewLess(left, right slang.Expression) slang.Expression {
	return &Less{
		BinaryOperation: NewBinaryOperation(slang.ExprLESS, left, right),
	}
}
