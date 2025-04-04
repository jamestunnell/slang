package expressions

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/types"
	"golang.org/x/exp/slices"
)

type Func struct {
	Inputs     []*types.NameType `json:"inputs"`
	Outputs    []*types.NameType `json:"outputs"`
	Statements []slang.Statement `json:"statements"`
}

func NewFunc(
	inputs, outputs []*types.NameType,
	statements ...slang.Statement,
) *Expression {
	return NewExpression(slang.ExprFUNC, &Func{
		Inputs:     inputs,
		Outputs:    outputs,
		Statements: statements,
	})
}

func (f *Func) IsEqual(other Core) bool {
	f2, ok := other.(*Func)
	if !ok {
		return false
	}

	if !slices.EqualFunc(f.Inputs, f2.Inputs, nameTypesEqual) {
		return false
	}

	if !slices.EqualFunc(f.Outputs, f2.Outputs, nameTypesEqual) {
		return false
	}

	if !slices.EqualFunc(f.Statements, f2.Statements, slang.StatementsEqual) {
		return false
	}

	return true
}

func nameTypesEqual(a, b *types.NameType) bool {
	return a.IsEqual(b)
}
