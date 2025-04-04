package statements

import (
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
)

type ForEach struct {
	Vars  []string         `json:"vars"`
	Expr  slang.Expression `json:"expr"`
	Block slang.Statement  `json:"block"`
}

func NewForEach(
	vars []string,
	expr slang.Expression,
	block slang.Statement,
) *Statement {
	core := &ForEach{Vars: vars, Expr: expr, Block: block}

	return NewStatement(slang.StatementFOREACH, core)
}

func (f *ForEach) IsEqual(other Core) bool {
	f2, ok := other.(*ForEach)
	if !ok {
		return false
	}

	if !slices.Equal(f.Vars, f2.Vars) {
		return false
	}

	if !f.Expr.IsEqual(f2.Expr) {
		return false
	}

	return f.Block.IsEqual(f2.Block)
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
