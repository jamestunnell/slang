package expressions

import "github.com/jamestunnell/slang"

func NewEqual(left, right *Expression) *Expression {
	return NewBinaryOperation(slang.ExprEQUAL, left, right)
}
