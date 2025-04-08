package expressions

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/field"
	"golang.org/x/exp/slices"
)

type Func struct {
	Inputs     []*field.Field    `json:"inputs"`
	Outputs    []*field.Field    `json:"outputs"`
	Statements []slang.Statement `json:"statements"`
}

func NewFunc(
	inputs, outputs []*field.Field,
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

	if !slices.EqualFunc(f.Inputs, f2.Inputs, fieldsEqual) {
		return false
	}

	if !slices.EqualFunc(f.Outputs, f2.Outputs, fieldsEqual) {
		return false
	}

	if !slices.EqualFunc(f.Statements, f2.Statements, slang.StatementsEqual) {
		return false
	}

	return true
}

func fieldsEqual(a, b *field.Field) bool {
	return a.IsEqual(b)
}
