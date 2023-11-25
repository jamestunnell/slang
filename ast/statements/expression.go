package statements

import (
	"github.com/jamestunnell/slang"
)

type Expression struct {
	*Base

	Value slang.Expression `json:"value"`
}

func NewExpression(val slang.Expression) *Expression {
	return &Expression{
		Base:  NewBase(slang.StatementEXPRESSION),
		Value: val,
	}
}

func (e *Expression) Equal(other slang.Statement) bool {
	e2, ok := other.(*Expression)
	if !ok {
		return false
	}

	return e2.Value.Equal(e.Value)
}

// func (st *Expression) Eval(env *slang.Environment) (slang.Object, error) {
// 	return st.Value.Eval(env)
// }
