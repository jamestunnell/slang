package statements

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
)

type Assign struct {
	Target *expressions.Expression `json:"target"`
	Value  *expressions.Expression `json:"value"`
}

func NewAssign(target, val *expressions.Expression) *Statement {
	core := &Assign{Target: target, Value: val}

	return NewStatement(slang.StatementASSIGN, core)
}

func (a *Assign) IsEqual(other Core) bool {
	a2, ok := other.(*Assign)
	if !ok {
		return false
	}

	return a.Target.IsEqual(a2.Target) && a2.Value.IsEqual(a.Value)
}

// func (st *Assign) Eval(env *slang.Environment) (slang.Object, error) {
// 	obj, err := st.Value.Eval(env)
// 	if err != nil {
// 		return objects.NULL(), err
// 	}

// 	env.Set(st.ident.Name, obj)

// 	return obj, nil
// }
