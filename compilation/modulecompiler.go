package compilation

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type ModuleCompiler struct {
	*Compiler
}

func NewModuleCompiler() *ModuleCompiler {
	return &ModuleCompiler{
		Compiler: NewCompiler(nil),
	}
}

func (c *ModuleCompiler) Run(stmts StmtSeq) bool {
	if !c.firstPass(stmts) {
		return false
	}

	return false
}

func (c *ModuleCompiler) firstPass(stmts StmtSeq) bool {
	stmt, ok := stmts.Next()
	for ok {
		switch stmt.Type() {
		case slang.StatementCLASS:
			name := stmt.(*statements.Class).Name

			c.symbols[name] = NewSymbol(name, SymbolCLASS)
		case slang.StatementFUNC:
			name := stmt.(*statements.Class).Name

			c.symbols[name] = NewSymbol(name, SymbolFUNC)
		case slang.StatementUSE:
			use := stmt.(*statements.Use)
			name := use.PathParts[len(use.PathParts)-1]

			c.symbols[name] = NewSymbol(name, SymbolMODULE)
		}

		stmt, ok = stmts.Next()
	}

	return true
}
