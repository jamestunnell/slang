package expressions

import "github.com/jamestunnell/slang"

func NewNotEqual(left, right *Expression) *Expression {
	return NewBinaryOperation(slang.ExprNOTEQUAL, left, right)
}
