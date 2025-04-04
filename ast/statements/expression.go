package statements

import (
	"github.com/jamestunnell/slang"
)

type Expression struct {
	Value slang.Expression `json:"value"`
}

func NewExpression(val slang.Expression) *Statement {
	core := &Expression{Value: val}

	return NewStatement(slang.StatementEXPRESSION, core)
}

func (e *Expression) IsEqual(other Core) bool {
	e2, ok := other.(*Expression)
	if !ok {
		return false
	}

	return e2.Value.Equal(e.Value)
}

// func (st *Expression) Eval(env *slang.Environment) (slang.Object, error) {
// 	return st.Value.Eval(env)
// }
