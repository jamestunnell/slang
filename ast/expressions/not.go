package expressions

import (
	"github.com/jamestunnell/slang"
)

type Not struct {
	*UnaryOperation
}

func NewNot(val slang.Expression) slang.Expression {
	return NewUnaryOperation(slang.ExprNOT, val)
}

// func (expr *Not) Eval(env *slang.Environment) (slang.Object, error) {
// 	obj, err := expr.Value.Eval(env)
// 	if err != nil {
// 		return objects.NULL(), err
// 	}

// 	return obj.Send(slang.MethodNOT)
// }
