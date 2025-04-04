package expressions

import "github.com/jamestunnell/slang"

func NewOr(left, right *Expression) *Expression {
	return NewBinaryOperation(slang.ExprOR, left, right)
}
