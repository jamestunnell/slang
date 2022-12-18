package statements

import (
	"github.com/jamestunnell/slang"
)

type Return struct {
	value slang.Expression
}

func NewReturn(value slang.Expression) *Return {
	return &Return{value: value}
}

func (r *Return) Type() slang.StatementType {
	return slang.StatementRETURN
}

func (r *Return) Equal(other slang.Statement) bool {
	r2, ok := other.(*Return)
	if !ok {
		return false
	}

	return r2.value.Equal(r.value)
}

func (st *Return) Eval(env *slang.Environment) (slang.Object, error) {
	return st.value.Eval(env)
}
