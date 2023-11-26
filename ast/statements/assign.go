package statements

import (
	"github.com/jamestunnell/slang"
)

type Assign struct {
	*Base

	Target slang.Expression `json:"target"`
	Value  slang.Expression `json:"value"`
}

func NewAssign(Target slang.Expression, val slang.Expression) slang.Statement {
	return &Assign{
		Base:   NewBase(slang.StatementASSIGN),
		Target: Target,
		Value:  val,
	}
}

func (a *Assign) Equal(other slang.Statement) bool {
	a2, ok := other.(*Assign)
	if !ok {
		return false
	}

	return a.Target.Equal(a2.Target) && a2.Value.Equal(a.Value)
}

// func (st *Assign) Eval(env *slang.Environment) (slang.Object, error) {
// 	obj, err := st.Value.Eval(env)
// 	if err != nil {
// 		return objects.NULL(), err
// 	}

// 	env.Set(st.ident.Name, obj)

// 	return obj, nil
// }
