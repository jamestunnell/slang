package compilation

import "github.com/jamestunnell/slang"

type MethodCompiler struct {
	*FuncCompiler
}

func NewMethodCompiler(symbol *slang.Symbol, stmts StmtSeq, parent Compiler) *MethodCompiler {
	return &MethodCompiler{
		FuncCompiler: NewFuncCompiler(symbol, stmts, parent),
	}
}
