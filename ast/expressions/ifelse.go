package expressions

import (
	"github.com/jamestunnell/slang"
)

type IfElse struct {
	*Base

	Condition    slang.Expression  `json:"condition"`
	Consequences []slang.Statement `json:"consequences"`
	Alternatives []slang.Statement `json:"alternatives"`
}

func NewIfElse(cond slang.Expression, conseqs, alterns []slang.Statement) *IfElse {
	return &IfElse{
		Base:         NewBase(slang.ExprIFELSE),
		Alternatives: alterns,
	}
}

func (i *IfElse) Equal(other slang.Expression) bool {
	i2, ok := other.(*IfElse)
	if !ok {
		return false
	}

	if !i.Condition.Equal(i2.Condition) {
		return false
	}

	if !statementsEqual(i.Consequences, i2.Consequences) {
		return false
	}

	return statementsEqual(i.Alternatives, i2.Alternatives)
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
