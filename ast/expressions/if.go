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

	return slang.StatementsEqual(i.Consequences, i2.Consequences)
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
