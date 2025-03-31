package expressions

import (
	"github.com/jamestunnell/slang"
)

type Func struct {
	*Base

	Params      []slang.Param     `json:"params"`
	ReturnTypes []slang.Type      `json:"returnTypes"`
	Statements  []slang.Statement `json:"statements"`
}

func NewFunc(
	params []slang.Param,
	returnTypes []slang.Type,
	statements ...slang.Statement,
) *Func {
	return &Func{
		Base:        NewBase(slang.ExprFUNC),
		Params:      params,
		ReturnTypes: returnTypes,
		Statements:  statements,
	}
}

func (f *Func) Equal(other slang.Expression) bool {
	f2, ok := other.(*Func)
	if !ok {
		return false
	}

	if len(f.Params) != len(f2.Params) {
		return false
	}

	for i, param := range f.Params {
		if param.GetName() != f2.Params[i].GetName() {
			return false
		}

		if !param.GetType().IsEqual(f2.Params[i].GetType()) {
			return false
		}
	}

	if len(f.ReturnTypes) != len(f2.ReturnTypes) {
		return false
	}

	for i, retType := range f.ReturnTypes {
		if !retType.IsEqual(f2.ReturnTypes[i]) {
			return false
		}
	}

	if len(f.Statements) != len(f2.Statements) {
		return false
	}

	for i, stmt := range f.Statements {
		if !stmt.Equal(f2.Statements[i]) {
			return false
		}
	}

	return true
}
