package expressions

import "github.com/jamestunnell/slang"

func NewLess(left, right slang.Expression) slang.Expression {
	return NewBinaryOperation(slang.ExprLESS, slang.MethodLT, left, right)
}
