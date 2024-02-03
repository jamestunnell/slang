package compilation

import "github.com/jamestunnell/slang"

type VarCompiler struct {
	*Base
}

func NewVarCompiler(symbol *slang.Symbol, parent Compiler) *VarCompiler {
	return &VarCompiler{
		Base: NewBase(symbol, NewEmptyStmtSeq(), parent),
	}
}

func (c *VarCompiler) FirstPass() error {
	return nil
}
