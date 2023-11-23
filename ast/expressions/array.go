package expressions

import (
	"github.com/jamestunnell/slang"
	"golang.org/x/exp/slices"
)

type Array struct {
	Elements []slang.Expression
}

func NewArray(elems ...slang.Expression) slang.Expression {
	return &Array{
		Elements: elems,
	}
}

func (c *Array) Type() slang.ExprType {
	return slang.ExprARRAY
}

func (c *Array) Equal(other slang.Expression) bool {
	c2, ok := other.(*Array)
	if !ok {
		return false
	}

	return slices.EqualFunc(c.Elements, c2.Elements, expressionsEqual)
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
