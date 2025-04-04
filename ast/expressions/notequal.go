package expressions

import "github.com/jamestunnell/slang"

func NewNotEqual(left, right slang.Expression) *Expression {
	return NewBinaryOperation(slang.ExprNOTEQUAL, left, right)
}
