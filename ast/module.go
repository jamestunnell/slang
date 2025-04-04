package ast

import (
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type Module struct {
	PathParts  []string                `json:"pathParts"`
	Statements []*statements.Statement `json:"statements"`
	// Errors    []error          `json:"errors"`
}

func NewModule(pathParts []string, stmts ...*statements.Statement) *Module {
	return &Module{
		PathParts:  pathParts,
		Statements: stmts,
		// Errors:    []error{},
	}
}

func (m *Module) GetPathParts() []string {
	return m.PathParts
}

func (m *Module) GetStructures() []slang.Structure {
	structs := []slang.Structure{}

	for _, stmt := range m.Statements {
		if stmt.GetType() == slang.StatementSTRUCT {
			if core, ok := stmt.Core.(*statements.Struct); ok {
				s := &Structure{
					Name:    core.Name,
					Comment: stmt.Comment,
					Fields:  core.Fields,
				}

				structs = append(structs, s)
			}
		}
	}

	return structs
}

func (m *Module) GetFunctions() []slang.Function {
	funcs := []slang.Function{}

	for _, stmt := range m.Statements {
		if stmt.GetType() == slang.StatementFUNC {
			if core, ok := stmt.Core.(*statements.Func); ok {
				f := &Function{
					Name:    core.Name,
					Comment: stmt.Comment,
					Inputs:  core.Inputs,
					Outputs: core.Outputs,
				}

				funcs = append(funcs, f)
			}
		}
	}

	return funcs
}

func (m *Module) IsEqual(other slang.Module) bool {
	m2, ok := other.(*Module)
	if !ok {
		return false
	}

	if !slices.Equal(m.PathParts, m2.PathParts) {
		return false
	}

	if !slices.EqualFunc(m.Statements, m2.Statements, statementsEqual) {
		return false
	}

	return true
}

func statementsEqual(a, b *statements.Statement) bool {
	return slang.StatementsEqual(a, b)
}
