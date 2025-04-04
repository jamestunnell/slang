package expressions

import "github.com/jamestunnell/slang"

func NewGreater(left, right *Expression) *Expression {
	return NewBinaryOperation(slang.ExprGREATER, left, right)
}
