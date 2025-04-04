package statements

import (
	"slices"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/types"
)

type Func struct {
	Name       string            `json:"name"`
	Inputs     []*types.NameType `json:"inputs"`
	Outputs    []*types.NameType `json:"outputs"`
	Statements []*Statement      `json:"statements"`
}

func NewFunc(
	name string,
	inputs, outputs []*types.NameType,
	statements ...*Statement,
) *Statement {
	core := &Func{
		Name:       name,
		Inputs:     inputs,
		Outputs:    outputs,
		Statements: statements,
	}

	return NewStatement(slang.StatementFUNC, core)
}

func (f *Func) IsEqual(other Core) bool {
	f2, ok := other.(*Func)
	if !ok {
		return false
	}

	if f.Name != f2.Name {
		return false
	}

	if !slices.EqualFunc(f.Inputs, f2.Inputs, nameTypesEqual) {
		return false
	}

	if !slices.EqualFunc(f.Outputs, f2.Outputs, nameTypesEqual) {
		return false
	}

	if !slices.EqualFunc(f.Statements, f2.Statements, statementsEqual) {
		return false
	}

	return true
}

func (f *Func) GetName() string {
	return f.Name
}

func (f *Func) GetInputs() []slang.Param {
	params := make([]slang.Field, len(f.Inputs))

	for i, param := range f.Inputs {
		params[i] = param
	}

	return params
}

func (f *Func) GetOutputs() []slang.Param {
	params := make([]slang.Field, len(f.Outputs))

	for i, param := range f.Outputs {
		params[i] = param
	}

	return params
}

func nameTypesEqual(a, b *types.NameType) bool {
	return a.IsEqual(b)
}
