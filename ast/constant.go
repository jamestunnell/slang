package ast

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/ast/statements"
)

type Constant struct {
	Name  string
	Value *expressions.Expression
}

func NewConstant(s *statements.Statement) *Constant {
	core, ok := s.Core.(*statements.Const)
	if !ok {
		return &Constant{
			Name:  "",
			Value: expressions.NewEmpty(),
		}
	}

	return &Constant{
		Name:  core.Name,
		Value: core.Value,
	}
}

func (c *Constant) GetName() string {
	return c.Name
}

func (c *Constant) GetValue() slang.Expression {
	return c.Value
}
