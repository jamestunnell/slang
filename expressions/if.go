package expressions

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/objects"
)

type If struct {
	Condition   slang.Expression
	Consequence slang.Statement
}

func NewIf(cond slang.Expression, conseq slang.Statement) *If {
	return &If{
		Condition:   cond,
		Consequence: conseq,
	}
}

func (i *If) Type() slang.ExprType { return slang.ExprIF }

func (i *If) Equal(other slang.Expression) bool {
	i2, ok := other.(*If)
	if !ok {
		return false
	}

	return i2.Condition.Equal(i.Condition) && i2.Consequence.Equal(i.Consequence)
}

func (expr *If) Eval(env *slang.Environment) (slang.Object, error) {
	res, err := expr.Condition.Eval(env)
	if err != nil {
		return objects.NULL(), err
	}

	if res.Truthy() {
		return expr.Consequence.Eval(slang.NewEnvironment(env))
	}

	return objects.NULL(), nil
}
