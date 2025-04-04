package expressions

import "github.com/jamestunnell/slang"

func NewAnd(left, right *Expression) *Expression {
	return NewBinaryOperation(slang.ExprAND, left, right)
}
