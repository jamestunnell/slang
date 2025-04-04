package expressions

import "github.com/jamestunnell/slang"

func NewLess(left, right slang.Expression) *Expression {
	return NewBinaryOperation(slang.ExprLESS, left, right)
}
