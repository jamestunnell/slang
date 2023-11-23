package expressions

import (
	"github.com/jamestunnell/slang"
)

type IfElse struct {
	*If

	Alternatives []slang.Statement
}

func NewIfElse(cond slang.Expression, conseqs, alterns []slang.Statement) *IfElse {
	return &IfElse{
		If:           NewIf(cond, conseqs),
		Alternatives: alterns,
	}
}

func (i *IfElse) Type() slang.ExprType { return slang.ExprIFELSE }

func (i *IfElse) Equal(other slang.Expression) bool {
	i2, ok := other.(*IfElse)
	if !ok {
		return false
	}

	if len(i.Alternatives) != len(i2.Alternatives) {
		return false
	}

	for idx, stmt := range i.Alternatives {
		if !stmt.Equal(i2.Alternatives[idx]) {
			return false
		}
	}

	return i.If.Equal(i2.If)
}

// func (expr *IfElse) Eval(env *slang.Environment) (slang.Object, error) {
// 	res, err := expr.Condition.Eval(env)
// 	if err != nil {
// 		return objects.NULL(), err
// 	}

// 	if res.Truthy() {
// 		return expr.Consequence.Eval(slang.NewEnvironment(env))
// 	}

// 	return expr.Alternative.Eval(slang.NewEnvironment(env))
// }
