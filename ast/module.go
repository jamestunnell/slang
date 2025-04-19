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

func (m *Module) GetStatementsWhere(keep func(*statements.Statement) bool) []*statements.Statement {
	return sliceutil.Where(m.Statements, keep)
}

func (m *Module) GetConstants() []slang.Constant {
	return sliceutil.MapWhere(m.Statements,
		func(s *statements.Statement) bool { return s.GetType() == slang.StatementCONST },
		func(s *statements.Statement) slang.Constant { return NewConstant(s) },
	)
}

func (m *Module) GetFunctions() []slang.Function {
	return sliceutil.MapWhere(m.Statements,
		func(s *statements.Statement) bool { return s.GetType() == slang.StatementFUNC },
		func(s *statements.Statement) slang.Function { return NewFunction(s) },
	)
}

func (m *Module) GetImports() []slang.Import {
	return sliceutil.MapWhere(m.Statements,
		func(s *statements.Statement) bool { return s.GetType() == slang.StatementUSE },
		func(s *statements.Statement) slang.Import { return NewImport(s) },
	)
}

func (m *Module) GetInterfaces() []slang.Interface {
	return sliceutil.MapWhere(m.Statements,
		func(s *statements.Statement) bool { return s.GetType() == slang.StatementINTERFACE },
		func(s *statements.Statement) slang.Interface { return NewInterface(s) },
	)
}

func (m *Module) GetStructures() []slang.Structure {
	return sliceutil.MapWhere(m.Statements,
		func(s *statements.Statement) bool { return s.GetType() == slang.StatementSTRUCT },
		func(s *statements.Statement) slang.Structure { return NewStructure(s) },
	)
}
func (m *Module) GetVariables() []slang.Variable {
	return sliceutil.MapWhere(m.Statements,
		func(s *statements.Statement) bool { return s.GetType() == slang.StatementVAR },
		func(s *statements.Statement) slang.Variable { return NewVariable(s) },
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
