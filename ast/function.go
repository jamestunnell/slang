package ast

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/ast/types"
)

type Function struct {
	Name    string
	Comment string
	Inputs  []*types.NameType
	Outputs []*types.NameType
}

func NewFunction(s *statements.Statement) slang.Function {
	core, ok := s.Core.(*statements.Func)
	if !ok {
		return &Function{
			Name:    "",
			Comment: "",
			Inputs:  []*types.NameType{},
			Outputs: []*types.NameType{},
		}
	}

	return &Function{
		Name:    core.Name,
		Comment: s.Comment,
		Inputs:  core.Inputs,
		Outputs: core.Outputs,
	}
}

func (f *Function) GetName() string {
	return f.Name
}

func (f *Function) GetComment() string {
	return f.Comment
}

func (f *Function) GetInputs() []slang.Param {
	params := make([]slang.Field, len(f.Inputs))

	for i, p := range f.Inputs {
		params[i] = p
	}

	return params
}

func (f *Function) GetOutputs() []slang.Param {
	params := make([]slang.Field, len(f.Outputs))

	for i, p := range f.Outputs {
		params[i] = p
	}

	return params
}
