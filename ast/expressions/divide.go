package expressions

import "github.com/jamestunnell/slang"

func NewDivide(left, right *Expression) *Expression {
	return NewBinaryOperation(slang.ExprDIVIDE, left, right)
}
