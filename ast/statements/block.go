package statements

import "github.com/jamestunnell/slang"

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

	return slang.StatementsEqual(b.Statements, b2.Statements)
}
