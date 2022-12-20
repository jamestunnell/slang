package statements

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/objects"
)

type Block struct {
	Statements []slang.Statement
}

func NewBlock(stmts ...slang.Statement) *Block {
	return &Block{Statements: stmts}
}

func (b *Block) Type() slang.StatementType {
	return slang.StatementEXPRESSION
}

func (b *Block) Equal(other slang.Statement) bool {
	b2, ok := other.(*Block)
	if !ok {
		return false
	}

	if len(b.Statements) != len(b2.Statements) {
		return false
	}

	for i, s := range b.Statements {
		if !b2.Statements[i].Equal(s) {
			return false
		}
	}

	return true
}

func (st *Block) Eval(env *slang.Environment) (slang.Object, error) {
	lastResult := slang.Object(objects.NULL())

	for _, s := range st.Statements {
		obj, err := s.Eval(env)
		if err != nil {
			return objects.NULL(), err
		}

		lastResult = obj

		if s.Type() == slang.StatementRETURN {
			break
		}
	}

	return lastResult, nil
}
