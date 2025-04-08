package ast

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/field"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/sliceutil"
)

type Interface struct {
	Name          string
	FunctionSpecs []slang.FunctionSpec
}

func NewInterface(s *statements.Statement) *Interface {
	core, ok := s.Core.(*statements.Interface)
	if !ok {
		return &Interface{
			Name:          "",
			FunctionSpecs: []slang.FunctionSpec{},
		}
	}

	return &Interface{
		Name:          core.Name,
		FunctionSpecs: sliceutil.Map(core.Specs, NewFunctionSpec),
	}
}

func NewFunctionSpec(spec *statements.FuncSpec) slang.FunctionSpec {
	return slang.FunctionSpec{
		Name:    spec.Name,
		Inputs:  sliceutil.Map(spec.Inputs, func(f *field.Field) slang.Param { return f }),
		Outputs: sliceutil.Map(spec.Outputs, func(f *field.Field) slang.Param { return f }),
	}
}

func (i *Interface) GetName() string {
	return i.Name
}

func (i *Interface) GetFunctionSpecs() []slang.FunctionSpec {
	return i.FunctionSpecs
}
