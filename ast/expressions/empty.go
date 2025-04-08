package expressions

import "github.com/jamestunnell/slang"

type Empty struct{}

func NewEmpty() *Expression {
	return NewExpression(slang.ExprEMPTY, &Empty{})
}

func (e *Empty) IsEqual(other Core) bool {
	_, ok := other.(*Empty)

	return ok
}
