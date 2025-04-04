package statements

import (
	"github.com/jamestunnell/slang"
)

type If struct {
	Condition slang.Expression `json:"condition"`
	Block     slang.Statement  `json:"block"`
}

func NewIf(
	cond slang.Expression,
	ifBlock slang.Statement,
) *Statement {
	core := &If{Condition: cond, Block: ifBlock}

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

	return i.Block.IsEqual(i2.Block)
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
