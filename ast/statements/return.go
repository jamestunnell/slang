package statements

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/tokens"
)

type Return struct{}

func NewReturn() *Statement {
	return NewStatement(slang.StatementRETURN, &Return{})
}

func (r *Return) GetName() (string, bool) {
	return "", false
}

func (r *Return) IsEqual(other Core) bool {
	_, ok := other.(*Return)

	return ok
}

func (r *Return) Render(level int, w slang.CodeWriter) {
	w.WriteString(tokens.StrRETURN)
}

// func (st *Return) Eval(env *slang.Environment) (slang.Object, error) {
// 	return st.value.Eval(env)
// }
