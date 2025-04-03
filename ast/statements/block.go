package statements

import (
	"github.com/jamestunnell/slang"
	"golang.org/x/exp/slices"
)

type Block struct {
	*Base

	Statements []slang.Statement
}

func NewBlock(stmts ...slang.Statement) *Block {
	return &Block{
		Base:       NewBase(slang.StatementBLOCK),
		Statements: stmts,
	}
}

func (b *Block) Equal(other slang.Statement) bool {
	b2, ok := other.(*Block)
	if !ok {
		return false
	}

	return slices.EqualFunc(b.Statements, b2.Statements, slang.StatementsEqual)
}
