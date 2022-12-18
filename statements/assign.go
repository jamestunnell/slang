package statements

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/expressions"
	"github.com/jamestunnell/slang/objects"
)

type Assign struct {
	ident *expressions.Identifier
	value slang.Expression
}

func NewAssign(ident *expressions.Identifier, val slang.Expression) slang.Statement {
	return &Assign{
		ident: ident,
		value: val,
	}
}

func (a *Assign) Type() slang.StatementType {
	return slang.StatementASSIGN
}

func (a *Assign) Equal(other slang.Statement) bool {
	a2, ok := other.(*Assign)
	if !ok {
		return false
	}

	if !a2.ident.Equal(a.ident) {
		return false
	}

	return a2.value.Equal(a.value)
}

func (st *Assign) Eval(env *slang.Environment) (slang.Object, error) {
	obj, err := st.value.Eval(env)
	if err != nil {
		return objects.NULL(), err
	}

	env.Set(st.ident.Name, obj)

	return obj, nil
}
