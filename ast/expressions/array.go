package expressions

import (
	"github.com/jamestunnell/slang"
	"golang.org/x/exp/slices"
)

type Array struct {
	*Base

	ValueType string             `json:"valueType"`
	Values    []slang.Expression `json:"values"`
}

func NewArray(valType string, vals ...slang.Expression) slang.Expression {
	return &Array{
		Base:      NewBase(slang.ExprARRAY),
		ValueType: valType,
		Values:    vals,
	}
}

func (a *Array) Equal(other slang.Expression) bool {
	a2, ok := other.(*Array)
	if !ok {
		return false
	}

	if a.ValueType != a2.ValueType {
		return false
	}

	return slices.EqualFunc(a.Values, a2.Values, expressionsEqual)
}

// func (expr *Array) Eval(env *slang.Environment) (slang.Object, error) {
// 	vals := make([]slang.Object, len(expr.Elements))
// 	for i := 0; i < len(expr.Elements); i++ {
// 		val, err := expr.Elements[i].Eval(env)
// 		if err != nil {
// 			return objects.NULL(), err
// 		}

// 		vals[i] = val
// 	}

// 	return objects.NewArray(vals...), nil
// }
