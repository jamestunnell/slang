package statements

import "github.com/jamestunnell/slang"

type Return struct{}

func NewReturn() *Statement {
	return NewStatement(slang.StatementRETURN, &Return{})
}

func (r *Return) IsEqual(other Core) bool {
	_, ok := other.(*Return)

	return ok
}

// func (st *Return) Eval(env *slang.Environment) (slang.Object, error) {
// 	return st.value.Eval(env)
// }
