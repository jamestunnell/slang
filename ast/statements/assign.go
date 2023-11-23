package statements

import (
	"github.com/jamestunnell/slang"
)

type Assign struct {
	Targets []string
	Value   slang.Expression
}

func NewAssign(val slang.Expression, Targets ...string) slang.Statement {
	return &Assign{
		Targets: Targets,
		Value:   val,
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

	if len(a.Targets) != len(a2.Targets) {
		return false
	}

	for i, tgt := range a.Targets {
		if tgt != a2.Targets[i] {
			return false
		}
	}

	return a2.Value.Equal(a.Value)
}

// func (st *Assign) Eval(env *slang.Environment) (slang.Object, error) {
// 	obj, err := st.Value.Eval(env)
// 	if err != nil {
// 		return objects.NULL(), err
// 	}

// 	env.Set(st.ident.Name, obj)

// 	return obj, nil
// }
