package expressions

import "github.com/jamestunnell/slang"

type Greater struct {
	*BinaryOperation
}

func NewGreater(left, right slang.Expression) slang.Expression {
	return &Greater{
		BinaryOperation: NewBinaryOperation(slang.ExprGREATER, left, right),
	}
}
