package statements

import (
	"github.com/jamestunnell/slang"
)

type ReturnVal struct {
	Value slang.Expression `json:"value"`
}

func NewReturnVal(value slang.Expression) *Statement {
	core := &ReturnVal{Value: value}

	return NewStatement(slang.StatementRETURNVAL, core)
}

func (r *ReturnVal) IsEqual(other Core) bool {
	r2, ok := other.(*ReturnVal)
	if !ok {
		return false
	}

	return r2.Value.Equal(r.Value)
}

// func (st *ReturnVal) Eval(env *slang.Environment) (slang.Object, error) {
// 	return st.value.Eval(env)
// }
