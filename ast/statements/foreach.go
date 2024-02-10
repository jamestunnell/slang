package statements

import (
	"github.com/jamestunnell/slang"
	"golang.org/x/exp/slices"
)

type ForEach struct {
	*Base

	Vars  []string         `json:"vars"`
	Expr  slang.Expression `json:"expr"`
	Block slang.Statement  `json:"block"`
}

func NewForEach(vars []string, expr slang.Expression, block slang.Statement) *ForEach {
	return &ForEach{
		Base:  NewBase(slang.StatementFOREACH),
		Vars:  vars,
		Expr:  expr,
		Block: block,
	}
}

func (f *ForEach) Equal(other slang.Statement) bool {
	f2, ok := other.(*ForEach)
	if !ok {
		return false
	}

	if !slices.Equal(f.Vars, f2.Vars) {
		return false
	}

	if !f.Expr.Equal(f2.Expr) {
		return false
	}

	return f.Block.Equal(f2.Block)
}

// func (expr *ForEach) Eval(env *slang.Environment) (slang.Object, error) {
// 	res, err := expr.Condition.Eval(env)
// 	if err != nil {
// 		return objects.NULL(), err
// 	}

// 	if res.Truthy() {
// 		return expr.Consequence.Eval(slang.NewEnvironment(env))
// 	}

// 	return objects.NULL(), nil
// }
