package statements

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
	"golang.org/x/exp/slices"
)

type If struct {
	Condition  *expressions.Expression `json:"condition"`
	Statements []*Statement            `json:"statements"`
}

func NewIf(
	cond *expressions.Expression,
	stmts []*Statement,
) *Statement {
	core := &If{
		Condition:  cond,
		Statements: stmts,
	}

	return NewStatement(slang.StatementIF, core)
}

func (i *If) IsEqual(other Core) bool {
	i2, ok := other.(*If)
	if !ok {
		return false
	}

	if !i.Condition.IsEqual(i2.Condition) {
		return false
	}

	return slices.EqualFunc(i.Statements, i2.Statements, statementsEqual)
}

// func (expr *If) Eval(env *slang.Environment) (slang.Object, error) {
// 	res, err := expr.Condition.Eval(env)
// 	if err != nil {
// 		return objects.NULL(), err
// 	}

// 	if res.Truthy() {
// 		return expr.Consequence.Eval(slang.NewEnvironment(env))
// 	}

// 	return objects.NULL(), nil
// }
