package expressions

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/objects"
)

type IfElse struct {
	Condition                slang.Expression
	Consequence, Alternative slang.Statement
}

func NewIfElse(cond slang.Expression, conseq, altern slang.Statement) *IfElse {
	return &IfElse{
		Condition:   cond,
		Consequence: conseq,
		Alternative: altern,
	}
}

func (i *IfElse) Type() slang.ExprType { return slang.ExprIF }

func (i *IfElse) Equal(other slang.Expression) bool {
	i2, ok := other.(*IfElse)
	if !ok {
		return false
	}

	return i2.Condition.Equal(i.Condition) &&
		i2.Consequence.Equal(i.Consequence) &&
		i2.Alternative.Equal(i.Alternative)
}

func (expr *IfElse) Eval(env *slang.Environment) (slang.Object, error) {
	res, err := expr.Condition.Eval(env)
	if err != nil {
		return objects.NULL(), err
	}

	if res.Truthy() {
		return expr.Consequence.Eval(slang.NewEnvironment(env))
	}

	return expr.Alternative.Eval(slang.NewEnvironment(env))
}
