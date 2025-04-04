package statements

import (
	"slices"

	"github.com/jamestunnell/slang"
)

type Func struct {
	Name       string        `json:"name"`
	Inputs     []slang.Param `json:"inputs"`
	Outputs    []slang.Param `json:"outputs"`
	Statements []*Statement  `json:"statements"`
}

func NewFunc(
	name string,
	inParams, outParams []slang.Param,
	statements ...*Statement,
) *Statement {
	core := &Func{
		Name:       name,
		Inputs:     inParams,
		Outputs:    outParams,
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

	if !slices.EqualFunc(f.Inputs, f2.Inputs, slang.NameTypesEqual) {
		return false
	}

	if !slices.EqualFunc(f.Outputs, f2.Outputs, slang.NameTypesEqual) {
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
	return f.Inputs
}

func (f *Func) GetOutputs() []slang.Param {
	return f.Outputs
}
