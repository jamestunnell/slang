package expressions

import "github.com/jamestunnell/slang"

func NewLessEqual(left, right *Expression) *Expression {
	return NewBinaryOperation(slang.ExprLESSEQUAL, left, right)
}
