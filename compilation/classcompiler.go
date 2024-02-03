package compilation

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/customerrs"
)

type ClassCompiler struct {
	*Base
}

func NewClassCompiler(symbol *slang.Symbol, stmts StmtSeq, parent Compiler) *ClassCompiler {
	return &ClassCompiler{
		Base: NewBase(symbol, stmts, parent),
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
			childSymbol := slang.NewChildSymbol(name, slang.SymbolFIELD, c.symbol)

			c.children[name] = NewFieldCompiler(childSymbol, c)
		case slang.StatementVAR:
			c.handleVarStmtFirstPass(stmt, c)
		case slang.StatementMETHOD:
			methStmt := stmt.(*statements.Method)
			name := methStmt.Name
			childSymbol := slang.NewChildSymbol(name, slang.SymbolMETHOD, c.symbol)
			childStmts := NewStmtSeq(methStmt.Function.Statements)

			c.children[name] = NewMethodCompiler(childSymbol, childStmts, c)
		default:
			return customerrs.NewErrTypeNotAllowed(
				stmt.Type().String()+" statement", "class body")
		}

		c.stmts.Advance()

		stmt, ok = c.stmts.Current()
	}

	return c.runChildFirstPasses()
}
