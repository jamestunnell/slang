package statements

import (
	"github.com/jamestunnell/slang"
)

type Return struct {
	Value slang.Expression `json:"value"`
}

func NewReturn(value slang.Expression) *Return {
	return &Return{Value: value}
}

func (r *Return) Type() slang.StatementType {
	return slang.StatementRETURN
}

func (r *Return) Equal(other slang.Statement) bool {
	r2, ok := other.(*Return)
	if !ok {
		return false
	}

	return r2.Value.Equal(r.Value)
}

// func (st *Return) Eval(env *slang.Environment) (slang.Object, error) {
// 	return st.value.Eval(env)
// }
