package ast

import (
	"github.com/jamestunnell/slang"
)

type Function struct {
	Comment     string            `json:"comment"`
	Params      []*Param          `json:"params"`
	ReturnTypes []string          `json:"returnTypes"`
	Body        []slang.Statement `json:"body"`
}

func NewFunction(params []*Param, returnTypes []string, body ...slang.Statement) *Function {
	return &Function{
		Params:      params,
		ReturnTypes: returnTypes,
		Body:        body,
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
		if !param.Equal(other.Params[i]) {
			return false
		}
	}

	if len(fn.ReturnTypes) != len(other.ReturnTypes) {
		return false
	}

	for i, retType := range fn.ReturnTypes {
		if retType != other.ReturnTypes[i] {
			return false
		}
	}

	if len(fn.Body) != len(other.Body) {
		return false
	}

	for i, stmt := range fn.Body {
		if !stmt.Equal(other.Body[i]) {
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
		names[i] = param.Name
	}

	return names
}

func (fn *Function) GetParamType(name string) (string, bool) {
	for _, param := range fn.Params {
		if name == param.Name {
			return param.Type, true
		}
	}

	return "", false
}

func (fn *Function) GetReturnTypes() []string {
	return fn.ReturnTypes
}
