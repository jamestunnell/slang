package ast

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/ast/statements"
)

type Variable struct {
	Name         string
	InitialValue *expressions.Expression
}

func NewVariable(s *statements.Statement) *Variable {
	core, ok := s.Core.(*statements.Var)
	if !ok {
		return &Variable{
			Name:         "",
			InitialValue: expressions.NewEmpty(),
		}
	}

	return &Variable{
		Name:         core.Name,
		InitialValue: core.InitialValue,
	}
}

func (v *Variable) GetName() string {
	return v.Name
}

func (v *Variable) GetInitialValue() slang.Expression {
	return v.InitialValue
}
