package statements

import (
	"github.com/jamestunnell/slang"
)

type ReturnVal struct {
	*Base

	Value slang.Expression `json:"value"`
}

func NewReturnVal(value slang.Expression) *ReturnVal {
	return &ReturnVal{
		Base:  NewBase(slang.StatementRETURNVAL),
		Value: value,
	}
}

func (r *ReturnVal) Equal(other slang.Statement) bool {
	r2, ok := other.(*ReturnVal)
	if !ok {
		return false
	}

	return r2.Value.Equal(r.Value)
}

// func (st *ReturnVal) Eval(env *slang.Environment) (slang.Object, error) {
// 	return st.value.Eval(env)
// }
