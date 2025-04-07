package ast

import (
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/sliceutil"
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
	return sliceutil.MapWhere(m.Statements,
		func(s *statements.Statement) bool { return s.GetType() == slang.StatementSTRUCT },
		NewStructure,
	)
}

func (m *Module) GetFunctions() []slang.Function {
	return sliceutil.MapWhere(m.Statements,
		func(s *statements.Statement) bool { return s.GetType() == slang.StatementFUNC },
		NewFunction,
	)
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
