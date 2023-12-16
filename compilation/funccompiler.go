package compilation

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/customerrs"
)

type FuncCompiler struct {
	*Base
}

func NewFuncCompiler(symbol *slang.Symbol, stmts StmtSeq, parent Compiler) *FuncCompiler {
	return &FuncCompiler{
		Base: NewBase(symbol, stmts, parent),
	}
}

func (c *FuncCompiler) FirstPass() error {
	stmt, ok := c.stmts.Current()

	for ok {
		switch stmt.Type() {
		case slang.StatementASSIGN,
			slang.StatementEXPRESSION,
			slang.StatementIF,
			slang.StatementIFELSE,
			slang.StatementRETURN,
			slang.StatementRETURNVAL,
			slang.StatementBLOCK:
			// do nothing on first pass
		default:
			return customerrs.NewErrTypeNotAllowed(
				stmt.Type().String()+" statement", "func body")
		}

		c.stmts.Advance()

		stmt, ok = c.stmts.Current()
	}

	return c.runChildFirstPasses()
}
