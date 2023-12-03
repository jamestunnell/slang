package expressions

import (
	"github.com/jamestunnell/slang"
)

type Negative struct {
	*UnaryOperation
}

func NewNegative(val slang.Expression) slang.Expression {
	return NewUnaryOperation(slang.ExprNEGATIVE, val)
}

// func (expr *Negative) Eval(env *slang.Environment) (slang.Object, error) {
// 	obj, err := expr.Value.Eval(env)
// 	if err != nil {
// 		return objects.NULL(), err
// 	}

// 	return obj.Send(slang.MethodNEG)
// }
