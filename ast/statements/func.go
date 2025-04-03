package statements

import (
	"slices"

	"github.com/jamestunnell/slang"
)

type Func struct {
	*Base

	Name       string            `json:"name"`
	InParams   []slang.Param     `json:"inputParams"`
	OutParams  []slang.Param     `json:"outputParams"`
	Statements []slang.Statement `json:"statements"`
}

func NewFunc(
	name string,
	inParams, outParams []slang.Param,
	statements ...slang.Statement,
) *Func {
	return &Func{
		Base:       NewBase(slang.StatementFUNC),
		Name:       name,
		InParams:   inParams,
		OutParams:  outParams,
		Statements: statements,
	}
}

func (f *Func) Equal(other slang.Statement) bool {
	f2, ok := other.(*Func)
	if !ok {
		return false
	}

	if f.Name != f2.Name {
		return false
	}

	if !slices.EqualFunc(f.InParams, f2.InParams, slang.NameTypesEqual) {
		return false
	}

	if !slices.EqualFunc(f.OutParams, f2.OutParams, slang.NameTypesEqual) {
		return false
	}

	if !slices.EqualFunc(f.Statements, f2.Statements, slang.StatementsEqual) {
		return false
	}

	return true
}

func (f *Func) GetName() string {
	return f.Name
}

func (f *Func) GetInputParams() []slang.Param {
	return f.InParams
}

func (f *Func) GetOutputParams() []slang.Param {
	return f.OutParams
}
