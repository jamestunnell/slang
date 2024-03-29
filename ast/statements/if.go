package statements

import (
	"github.com/jamestunnell/slang"
)

type If struct {
	*Base

	Condition slang.Expression `json:"condition"`
	Block     slang.Statement  `json:"block"`
}

func NewIf(cond slang.Expression, ifBlock slang.Statement) *If {
	return &If{
		Base:      NewBase(slang.StatementIF),
		Condition: cond,
		Block:     ifBlock,
	}
}

func (i *If) Equal(other slang.Statement) bool {
	i2, ok := other.(*If)
	if !ok {
		return false
	}

	if !i.Condition.Equal(i2.Condition) {
		return false
	}

	return i.Block.Equal(i2.Block)
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
