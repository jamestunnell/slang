package expressions

import (
	"github.com/jamestunnell/slang"
	"golang.org/x/exp/slices"
)

type Func struct {
	*Base

	Inputs     []slang.Param     `json:"inputs"`
	Outputs    []slang.Param     `json:"outputs"`
	Statements []slang.Statement `json:"statements"`
}

func NewFunc(
	inParams, outParams []slang.Param,
	statements ...slang.Statement,
) *Func {
	return &Func{
		Base:       NewBase(slang.ExprFUNC),
		Inputs:     inParams,
		Outputs:    outParams,
		Statements: statements,
	}
}

func (f *Func) Equal(other slang.Expression) bool {
	f2, ok := other.(*Func)
	if !ok {
		return false
	}

	if !slices.EqualFunc(f.Inputs, f2.Inputs, slang.NameTypesEqual) {
		return false
	}

	if !slices.EqualFunc(f.Outputs, f2.Outputs, slang.NameTypesEqual) {
		return false
	}

	if !slices.EqualFunc(f.Statements, f2.Statements, slang.StatementsEqual) {
		return false
	}

	return true
}
