package expressions

import (
	"github.com/jamestunnell/slang"
	"golang.org/x/exp/slices"
)

type Func struct {
	Inputs     []slang.Param     `json:"inputs"`
	Outputs    []slang.Param     `json:"outputs"`
	Statements []slang.Statement `json:"statements"`
}

func NewFunc(
	inParams, outParams []slang.Param,
	statements ...slang.Statement,
) *Expression {
	return NewExpression(slang.ExprFUNC, &Func{
		Inputs:     inParams,
		Outputs:    outParams,
		Statements: statements,
	})
}

func (f *Func) IsEqual(other Core) bool {
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
