package statements

import (
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
)

type ForEach struct {
	Vars       []string                `json:"vars"`
	Expr       *expressions.Expression `json:"expr"`
	Statements []*Statement            `json:"statements"`
}

func NewForEach(
	vars []string,
	expr *expressions.Expression,
	stmts []*Statement,
) *Statement {
	core := &ForEach{Vars: vars, Expr: expr, Statements: stmts}

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

	return slices.EqualFunc(f.Statements, f2.Statements, statementsEqual)
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
