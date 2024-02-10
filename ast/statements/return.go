package statements

import (
	"github.com/jamestunnell/slang"
)

type Return struct {
	*Base
}

func NewReturn() *Return {
	return &Return{
		Base: NewBase(slang.StatementRETURN),
	}
}

func (r *Return) Equal(other slang.Statement) bool {
	_, ok := other.(*Return)

	return ok
}

// func (st *Return) Eval(env *slang.Environment) (slang.Object, error) {
// 	return st.value.Eval(env)
// }
