package statements

import (
	"github.com/jamestunnell/slang"
)

type Break struct {
	*Base
}

func NewBreak() *Break {
	return &Break{
		Base: NewBase(slang.StatementBREAK),
	}
}

func (f *Break) Equal(other slang.Statement) bool {
	_, ok := other.(*Break)
	return ok
}

// func (expr *Break) Eval(env *slang.Environment) (slang.Object, error) {
// 	res, err := expr.Condition.Eval(env)
// 	if err != nil {
// 		return objects.NULL(), err
// 	}

// 	if res.Truthy() {
// 		return expr.Consequence.Eval(slang.NewEnvironment(env))
// 	}

// 	return objects.NULL(), nil
// }
