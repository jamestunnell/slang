package expressions

import (
	"github.com/jamestunnell/slang"
)

type If struct {
	*Base

	Condition    slang.Expression  `json:"condition"`
	Consequences []slang.Statement `json:"consequences"`
}

func NewIf(cond slang.Expression, conseqs []slang.Statement) *If {
	return &If{
		Base:         NewBase(slang.ExprIF),
		Condition:    cond,
		Consequences: conseqs,
	}
}

func (i *If) Equal(other slang.Expression) bool {
	i2, ok := other.(*If)
	if !ok {
		return false
	}

	if !i.Condition.Equal(i2.Condition) {
		return false
	}

	return statementsEqual(i.Consequences, i2.Consequences)
}

func statementsEqual(a, b []slang.Statement) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, stmt := range a {
		if !stmt.Equal(b[idx]) {
			return false
		}
	}

	return true
}

// func (expr *If) Eval(env *slang.Environment) (slang.Object, error) {
// 	res, err := expr.Condition.Eval(env)
// 	if err != nil {
// 		return objects.NULL(), err
// 	}

// 	if res.Truthy() {
// 		return expr.Consequence.Eval(slang.NewEnvironment(env))
// 	}

// 	return objects.NULL(), nil
// }
