package expressions

import "github.com/jamestunnell/slang"

func NewNotEqual(left, right slang.Expression) slang.Expression {
	return NewBinaryOperation(slang.ExprNOTEQUAL, left, right)
}
