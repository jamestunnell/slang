package expressions

import (
	"github.com/jamestunnell/slang"
)

type If struct {
	Condition    slang.Expression
	Consequences []slang.Statement
}

func NewIf(cond slang.Expression, conseqs []slang.Statement) *If {
	return &If{
		Condition:    cond,
		Consequences: conseqs,
	}
}

func (i *If) Type() slang.ExprType { return slang.ExprIF }

func (i *If) Equal(other slang.Expression) bool {
	i2, ok := other.(*If)
	if !ok {
		return false
	}

	if len(i.Consequences) != len(i2.Consequences) {
		return false
	}

	for idx, stmt := range i.Consequences {
		if !stmt.Equal(i2.Consequences[idx]) {
			return false
		}
	}

	return i2.Condition.Equal(i.Condition)
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
