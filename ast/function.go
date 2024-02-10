package ast

import (
	"github.com/jamestunnell/slang"
)

type Function struct {
	Comment     string            `json:"comment"`
	Params      []slang.Param     `json:"params"`
	ReturnTypes []slang.Type      `json:"returnTypes"`
	Statements  []slang.Statement `json:"statements"`
}

func NewFunction(params []slang.Param, returnTypes []slang.Type, body ...slang.Statement) *Function {
	return &Function{
		Params:      params,
		ReturnTypes: returnTypes,
		Statements:  body,
	}
}

func (fn *Function) Equal(other *Function) bool {
	if fn.Comment != other.Comment {
		return false
	}

	if len(fn.Params) != len(other.Params) {
		return false
	}

	for i, param := range fn.Params {
		if param.GetName() != other.Params[i].GetName() {
			return false
		}

		if !param.GetType().IsEqual(other.Params[i].GetType()) {
			return false
		}
	}

	if len(fn.ReturnTypes) != len(other.ReturnTypes) {
		return false
	}

	for i, retType := range fn.ReturnTypes {
		if !retType.IsEqual(other.ReturnTypes[i]) {
			return false
		}
	}

	if len(fn.Statements) != len(other.Statements) {
		return false
	}

	for i, stmt := range fn.Statements {
		if !stmt.Equal(other.Statements[i]) {
			return false
		}
	}

	return true
}

func (fn *Function) GetComment() string {
	return fn.Comment
}

func (fn *Function) GetParamNames() []string {
	names := make([]string, len(fn.Params))

	for i, param := range fn.Params {
		names[i] = param.GetName()
	}

	return names
}

func (fn *Function) GetParamType(name string) (slang.Type, bool) {
	for _, param := range fn.Params {
		if name == param.GetName() {
			return param.GetType(), true
		}
	}

	return nil, false
}

func (fn *Function) GetReturnTypes() []slang.Type {
	return fn.ReturnTypes
}
