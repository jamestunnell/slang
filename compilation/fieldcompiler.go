package compilation

import "github.com/jamestunnell/slang"

type FieldCompiler struct {
	*Base
}

func NewFieldCompiler(symbol *slang.Symbol, parent Compiler) *FieldCompiler {
	return &FieldCompiler{
		Base: NewBase(symbol, NewEmptyStmtSeq(), parent),
	}
}

func (c *FieldCompiler) FirstPass() error {
	return nil
}
