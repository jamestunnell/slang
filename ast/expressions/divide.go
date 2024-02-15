package expressions

import "github.com/jamestunnell/slang"

func NewDivide(left, right slang.Expression) slang.Expression {
	return NewBinaryOperation(slang.ExprDIVIDE, slang.MethodDIV, left, right)
}
