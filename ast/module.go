package ast

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type Module struct {
	Path       string                  `json:"path"`
	Statements []*statements.Statement `json:"statements"`
	Structures []slang.Structure       `json:"structures"`
	Functions  []slang.Function        `json:"functions"`
	// Errors    []error          `json:"errors"`
}

func NewModule(path string, stmts ...*statements.Statement) *Module {
	structures := []slang.Structure{}
	functions := []slang.Function{}

	for _, stmt := range stmts {
		if stmt.GetType() == slang.StatementSTRUCT {
			if core, ok := stmt.Core.(*statements.Struct); ok {
				s := &Structure{
					Name:    core.Name,
					Comment: stmt.Comment,
					Fields:  core.Fields,
				}

				structures = append(structures, s)
			}
		}

		if stmt.GetType() == slang.StatementFUNC {
			if core, ok := stmt.Core.(*statements.Func); ok {
				f := &Function{
					Name:    core.Name,
					Comment: stmt.Comment,
					Inputs:  core.Inputs,
					Outputs: core.Outputs,
				}

				functions = append(functions, f)
			}
		}
	}

	return &Module{
		Path:       path,
		Statements: stmts,
		Structures: structures,
		Functions:  functions,
		// Errors:    []error{},
	}
}

func (m *Module) GetPath() string {
	return m.Path
}

func (m *Module) GetStructures() []slang.Structure {
	return m.Structures
}

func (m *Module) GetFunctions() []slang.Function {
	return m.Functions
}
