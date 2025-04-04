package statements

import "github.com/jamestunnell/slang"

type Continue struct {
}

func NewContinue() *Statement {
	return NewStatement(slang.StatementCONTINUE, &Continue{})
}

func (f *Continue) IsEqual(other Core) bool {
	_, ok := other.(*Continue)

	return ok
}

// func (expr *Continue) Eval(env *slang.Environment) (slang.Object, error) {
// 	res, err := expr.Condition.Eval(env)
// 	if err != nil {
// 		return objects.NULL(), err
// 	}

// 	if res.Truthy() {
// 		return expr.Consequence.Eval(slang.NewEnvironment(env))
// 	}

// 	return objects.NULL(), nil
// }
