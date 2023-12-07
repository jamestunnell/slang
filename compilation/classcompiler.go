package compilation

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/customerrs"
)

type ClassCompiler struct {
	*Base
}

func NewClassCompiler(scope string, stmts StmtSeq, parent Compiler) *ClassCompiler {
	return &ClassCompiler{
		Base: NewBase(scope, stmts, parent),
	}
}

func (c *ClassCompiler) FirstPass() error {
	stmt, ok := c.stmts.Current()

	for ok {
		switch stmt.Type() {
		case slang.StatementCLASS:
			c.handleClassStmtFirstPass(stmt, c)
		case slang.StatementFUNC:
			c.handleFuncStmtFirstPass(stmt, c)
		case slang.StatementFIELD:
			fieldStmt := stmt.(*statements.Field)
			name := fieldStmt.Name

			c.symbols[name] = NewSymbol(c.scope, name, SymbolFIELD)
		case slang.StatementMETHOD:
			methStmt := stmt.(*statements.Method)
			name := methStmt.Name
			childStmts := NewStmtSeq(methStmt.Function.Statements)
			child := NewFuncCompiler(c.scopedName(name), childStmts, c)

			c.symbols[name] = NewSymbol(c.scope, name, SymbolMETHOD)
			c.children[name] = child
		default:
			return customerrs.NewErrTypeNotAllowed(
				stmt.Type().String()+" statement", "class body")
		}

		c.stmts.Advance()

		stmt, ok = c.stmts.Current()
	}

	return c.runChildFirstPasses()
}
