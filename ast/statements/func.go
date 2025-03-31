package statements

import (
	"github.com/jamestunnell/slang"
)

type Func struct {
	*Base

	Name        string            `json:"name"`
	Params      []slang.Param     `json:"params"`
	ReturnTypes []slang.Type      `json:"returnTypes"`
	Statements  []slang.Statement `json:"statements"`
}

func NewFunc(
	name string,
	params []slang.Param,
	returnTypes []slang.Type,
	statements ...slang.Statement,
) *Func {
	return &Func{
		Base:        NewBase(slang.StatementFUNC),
		Name:        name,
		Params:      params,
		ReturnTypes: returnTypes,
		Statements:  statements,
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

func (f *Func) GetName() string {
	return f.Name
}

func (f *Func) GetParamNames() []string {
	names := make([]string, len(f.Params))

	for i, param := range f.Params {
		names[i] = param.GetName()
	}

	return names
}

func (f *Func) GetParamType(name string) (slang.Type, bool) {
	for _, param := range f.Params {
		if name == param.GetName() {
			return param.GetType(), true
		}
	}

	return nil, false
}

func (f *Func) GetReturnTypes() []slang.Type {
	return f.ReturnTypes
}
