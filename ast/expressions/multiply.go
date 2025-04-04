package expressions

import "github.com/jamestunnell/slang"

func NewMultiply(left, right *Expression) *Expression {
	return NewBinaryOperation(slang.ExprMULTIPLY, left, right)
}
