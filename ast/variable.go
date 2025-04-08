package ast

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/ast/types"
)

type Variable struct {
	Name string
	Type *types.Type
}

func NewVariable(s *statements.Statement) *Variable {
	core, ok := s.Core.(*statements.Var)
	if !ok {
		return &Variable{
			Name: "",
			Type: types.NewEmpty(),
		}
	}

	return &Variable{
		Name: core.Name,
		Type: core.ValueType,
	}
}

func (v *Variable) GetName() string {
	return v.Name
}

func (v *Variable) GetType() slang.Type {
	return v.Type
}
