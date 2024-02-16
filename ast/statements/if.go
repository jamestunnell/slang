package statements

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/objects"
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

func (i *If) IsEqual(other slang.Statement) bool {
	i2, ok := other.(*If)
	if !ok {
		return false
	}

	if !i.Condition.IsEqual(i2.Condition) {
		return false
	}

	return i.Block.IsEqual(i2.Block)
}

func (expr *If) Eval(env slang.Environment) (slang.Objects, error) {
	res, err := expr.Condition.Eval(env)
	if err != nil {
		return nil, err
	}

	if res.IsEqual(objects.TRUE()) {
		return expr.Block.Eval(slang.NewEnv(env))
	}

	return slang.Objects{objects.NULL()}, nil
}
