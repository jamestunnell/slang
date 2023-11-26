package statements

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
)

type Method struct {
	*Base

	Name     string        `json:"name"`
	Function *ast.Function `json:"function"`
}

func NewMethod(name string, f *ast.Function) *Method {
	return &Method{
		Base:     NewBase(slang.StatementMETHOD),
		Name:     name,
		Function: f,
	}
}

func (m *Method) Equal(other slang.Statement) bool {
	m2, ok := other.(*Method)
	if !ok {
		return false
	}

	return m.Name == m2.Name && m.Function.Equal(m2.Function)
}
