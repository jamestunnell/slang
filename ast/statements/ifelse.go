package statements

import (
	"github.com/jamestunnell/slang"
)

type IfElse struct {
	Condition slang.Expression `json:"condition"`
	IfBlock   slang.Statement  `json:"ifBlock"`
	ElseBlock slang.Statement  `json:"elseBlock"`
}

func NewIfElse(
	cond slang.Expression,
	ifBlock, elseBlock slang.Statement,
) *Statement {
	core := &IfElse{
		Condition: cond,
		IfBlock:   ifBlock,
		ElseBlock: elseBlock,
	}

	return NewStatement(slang.StatementIFELSE, core)
}

func (i *IfElse) IsEqual(other Core) bool {
	i2, ok := other.(*IfElse)
	if !ok {
		return false
	}

	if !i.Condition.Equal(i2.Condition) {
		return false
	}

	if !i.IfBlock.IsEqual(i2.IfBlock) {
		return false
	}

	return i.ElseBlock.IsEqual(i2.ElseBlock)
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
