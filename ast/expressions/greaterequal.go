package expressions

import "github.com/jamestunnell/slang"

func NewGreaterEqual(left, right *Expression) *Expression {
	return NewBinaryOperation(slang.ExprGREATEREQUAL, left, right)
}
