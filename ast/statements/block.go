package statements

import (
	"github.com/jamestunnell/slang"
	"golang.org/x/exp/slices"
)

type Block struct {
	Statements []*Statement `json:"statements"`
}

func NewBlock(stmts ...*Statement) *Statement {
	core := &Block{Statements: stmts}

	return NewStatement(slang.StatementBLOCK, core)
}

func (b *Block) IsEqual(other Core) bool {
	b2, ok := other.(*Block)
	if !ok {
		return false
	}

	return slices.EqualFunc(b.Statements, b2.Statements, statementsEqual)
}

func statementsEqual(a, b *Statement) bool {
	return slang.StatementsEqual(a, b)
}
