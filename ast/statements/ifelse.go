package statements

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
	"golang.org/x/exp/slices"
)

type IfElse struct {
	Condition      *expressions.Expression `json:"condition"`
	IfStatements   []*Statement            `json:"ifStatements"`
	ElseStatements []*Statement            `json:"elseStatements"`
}

func NewIfElse(
	cond *expressions.Expression,
	ifStmts, elseStmts []*Statement,
) *Statement {
	core := &IfElse{
		Condition:      cond,
		IfStatements:   ifStmts,
		ElseStatements: elseStmts,
	}

	return NewStatement(slang.StatementIFELSE, core)
}

func (i *IfElse) IsEqual(other Core) bool {
	i2, ok := other.(*IfElse)
	if !ok {
		return false
	}

	if !i.Condition.IsEqual(i2.Condition) {
		return false
	}

	if !slices.EqualFunc(i.IfStatements, i2.IfStatements, statementsEqual) {
		return false
	}

	return slices.EqualFunc(i.ElseStatements, i2.ElseStatements, statementsEqual)
}

// func (expr *IfElse) Eval(env *slang.Environment) (slang.Object, error) {
// 	res, err := expr.Condition.Eval(env)
// 	if err != nil {
// 		return objects.NULL(), err
// 	}

// 	if res.Truthy() {
// 		return expr.Consequence.Eval(slang.NewEnvironment(env))
// 	}

// 	return expr.Alternative.Eval(slang.NewEnvironment(env))
// }
