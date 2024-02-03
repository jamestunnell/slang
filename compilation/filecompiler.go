package compilation

import (
	"strings"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/customerrs"
)

type FileCompiler struct {
	*Base

	Imports map[string]string
	Vars    []string
}

func NewFileCompiler(symbol *slang.Symbol, stmts StmtSeq) *FileCompiler {
	return &FileCompiler{
		Base:    NewBase(symbol, stmts, nil),
		Vars:    []string{},
		Imports: map[string]string{},
	}
}

func (c *FileCompiler) FirstPass() error {
	stmt, ok := c.stmts.Current()

	for ok {
		switch stmt.Type() {
		case slang.StatementCLASS:
			c.handleClassStmtFirstPass(stmt, c)
		case slang.StatementFUNC:
			c.handleFuncStmtFirstPass(stmt, c)
		case slang.StatementUSE:
			use := stmt.(*statements.Use)
			name := use.PathParts[len(use.PathParts)-1]

			c.Imports[name] = strings.Join(use.PathParts, "/")
		case slang.StatementVAR:
			c.handleVarStmtFirstPass(stmt, c)
		default:
			return customerrs.NewErrTypeNotAllowed(
				stmt.Type().String()+" statement", "module file")
		}

		c.stmts.Advance()

		stmt, ok = c.stmts.Current()
	}

	return c.runChildFirstPasses()
}
