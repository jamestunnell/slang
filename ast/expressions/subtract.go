package expressions

import "github.com/jamestunnell/slang"

func NewSubtract(left, right *Expression) *Expression {
	return NewBinaryOperation(slang.ExprSUBTRACT, left, right)
}
