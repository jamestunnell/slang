package expressions

import (
	"github.com/jamestunnell/slang"
	"golang.org/x/exp/slices"
)

type Func struct {
	*Base

	InParams   []slang.Param     `json:"inputParams"`
	OutParams  []slang.Param     `json:"outputParams"`
	Statements []slang.Statement `json:"statements"`
}

func NewFunc(
	inParams, outParams []slang.Param,
	statements ...slang.Statement,
) *Func {
	return &Func{
		Base:       NewBase(slang.ExprFUNC),
		InParams:   inParams,
		OutParams:  outParams,
		Statements: statements,
	}
}

func (f *Func) Equal(other slang.Expression) bool {
	f2, ok := other.(*Func)
	if !ok {
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
