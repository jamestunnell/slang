package statements

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/objects"
)

type IfElse struct {
	*Base

	Condition slang.Expression `json:"condition"`
	IfBlock   slang.Statement  `json:"ifBlock"`
	ElseBlock slang.Statement  `json:"elseBlock"`
}

func NewIfElse(cond slang.Expression, ifBlock, elseBlock slang.Statement) *IfElse {
	return &IfElse{
		Base:      NewBase(slang.StatementIFELSE),
		Condition: cond,
		IfBlock:   ifBlock,
		ElseBlock: elseBlock,
	}
}

func (i *IfElse) IsEqual(other slang.Statement) bool {
	i2, ok := other.(*IfElse)
	if !ok {
		return false
	}

	if !i.Condition.IsEqual(i2.Condition) {
		return false
	}

	if !i.IfBlock.IsEqual(i2.IfBlock) {
		return false
	}

	return i.ElseBlock.IsEqual(i2.ElseBlock)
}

func (expr *IfElse) Eval(env slang.Environment) (slang.Objects, error) {
	res, err := expr.Condition.Eval(env)
	if err != nil {
		return nil, err
	}

	if res.IsEqual(objects.TRUE()) {
		return expr.IfBlock.Eval(slang.NewEnv(env))
	}

	return expr.ElseBlock.Eval(slang.NewEnv(env))
}
